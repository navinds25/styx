package setup

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/nodeconfig"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func nodeConfigDB(dataRoot string) error {
	nodeConfigDBPath := filepath.Join(dataRoot, "node_config")
	if err := createDataDir(nodeConfigDBPath); err != nil {
		return err
	}
	encryptionKey, err := base64.StdEncoding.DecodeString(app.MainFlagVal.EncryptionKey)
	if err != nil {
		log.Error("Invalid encryption key", err)
		return err
	}
	dbOpts := badger.DefaultOptions(nodeConfigDBPath).WithEncryptionKey(encryptionKey).WithLogger(nil)
	nodeConfigDB, err := badger.Open(dbOpts)
	if err != nil {
		return err
	}
	nodeconfig.InitNodeConfigDB(nodeconfig.BadgerDB{
		NodeConfigDB: nodeConfigDB,
	})
	return nil
}

func createDataDir(datadir string) error {
	log.Debug("datadir:", datadir)
	_, err := os.Stat(datadir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(datadir, 0755); err != nil {
			return err
		}
		log.Info("created data directory: ", datadir)
	}
	return nil
}

// logic for this is weird but correct as it accounts for default value in Cli.
func readHostConfigFromCli() (*nodeconfig.HostConfigInput, error) {
	if app.MainFlagVal.Config != "" {
		hcInput, err := nodeconfig.HostConfigFromYAMLFile(app.MainFlagVal.Config)
		if err != nil {
			return nil, err
		}
		return hcInput, nil
	}
	return nil, nil
}

func readConfigFromCli() (*nodeconfig.ConfigInput, error) {
	if app.MainFlagVal.Config != "" {
		input, err := nodeconfig.ConfigFromYAMLFile(app.MainFlagVal.Config)
		if err != nil {
			return nil, err
		}
		return input, nil
	}
	return nil, nil
}

func overwriteFromCli(cliHC *nodeconfig.HostConfigInput) (*nodeconfig.HostConfigModel, error) {
	if cliHC == nil {
		return nil, errors.New("config didn't have any value")
	}
	hcM, err := nodeconfig.HostConfigToModel(cliHC)
	if err != nil {
		log.Debug("conversion of cli input to hostconfig model failed with ", err)
		return nil, err
	}
	if err := nodeconfig.Data.NodeConfig.AddHostConfigEntry(nodeconfig.HostConfigKey, hcM); err != nil {
		log.Debug("failed to add hostconfig entry to database ", err)
		return nil, err
	}
	log.Debug("added entry")
	dHcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
	if err != nil {
		log.Debug("failed to get hostconfig entry from database ", err)
		return dHcM, err
	}
	log.Debug("got entry from DB", dHcM)
	return dHcM, nil
}

func addNode(mc *nodeconfig.MasterConfigInput) error {
	if mc.MasterAddress == "" && mc.MasterIP != "" {
		mc.MasterAddress = mc.MasterIP + ":" + strconv.Itoa(mc.GRPCPort)
	}
	conn, err := grpc.Dial(mc.MasterAddress, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		log.Error("could not send request to master")
	}
	defer conn.Close()
	if err := nodeconfig.AddNodeClient(conn); err != nil {
		log.Error("error from AddNodeClient", err)
	}

	return nil
}

func updateHostConfig() (*nodeconfig.HostConfigModel, error) {
	if !app.MainFlagVal.OverwriteHostConfig {
		hcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
		if err != nil {
			log.Error("error reading hostconfig: ", err)
			return nil, err
		}
		log.Debugf("got hostconfig from db: %v", hcM)
		return hcM, nil
	}
	log.Debug("reading config from cli")
	config, err := readConfigFromCli()
	if err != nil {
		log.Error("error reading config", err)
		return nil, err
	}
	log.Debugf("got config: %+v", config.HostConfig)
	hcM, err := overwriteFromCli(config.HostConfig)
	if err != nil {
		log.Error("error updating hostconfig: ", err)
		return nil, err
	}
	log.Debugf("got hostconfig from db/overwrite: %v", hcM)
	if err := addNode(config.MasterConfig); err != nil {
		log.Error("error adding node to cluster", err)
		return nil, err
	}
	return hcM, nil
}

// NodeSetup is the main setup function
func NodeSetup() (*nodeconfig.HostConfigModel, error) {
	// read cli flags
	if err := app.MainFlagVal.CliSetDefaults(); err != nil {
		return nil, err
	}

	// setup dbs
	if err := nodeConfigDB(app.MainFlagVal.DataDir); err != nil {
		return nil, err
	}
	log.Debug("completed the db setup")
	// get the hostconfig
	hcM, err := updateHostConfig()
	if err != nil {
		log.Error("error updating hostconfig: ", err)
		return nil, err
	}
	return hcM, nil
}
