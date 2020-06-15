package setup

import (
	"crypto/x509"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

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
	dbOpts := badger.DefaultOptions(nodeConfigDBPath).WithTruncate(true).WithEncryptionKey(encryptionKey).WithLogger(nil)
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
	hcM, masterModel, err := nodeconfig.HostConfigToModel(cliHC)
	if err != nil {
		log.Debug("conversion of cli input to hostconfig model failed with ", err)
		return nil, err
	}
	if err := nodeconfig.Data.NodeConfig.AddHostConfigEntry(nodeconfig.HostConfigKey, hcM); err != nil {
		log.Debug("failed to add hostconfig entry to database ", err)
		return nil, err
	}
	log.Debug("added entry for hostconfig")
	if err := nodeconfig.Data.NodeConfig.AddMasterConfigEntry(masterModel); err != nil {
		log.Debug("failed to add entry for master config")
		return nil, err
	}
	dHcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
	if err != nil {
		log.Debug("failed to get hostconfig entry from database ", err)
		return dHcM, err
	}
	log.Debug("got entry from DB", dHcM)
	return dHcM, nil
}

func addNode(certFile string) error {
	mc, err := nodeconfig.Data.NodeConfig.GetMasterConfigEntry()
	if err != nil {
		return err
	}
	certPool := x509.NewCertPool()
	certFD, err := os.Open(certFile)
	if err != nil {
		log.Error("error opening tlscert", err)
		return err
	}
	cert, err := ioutil.ReadAll(certFD)
	if err != nil {
		log.Error("error reading tlscert: ", err)
		return err
	}
	status := certPool.AppendCertsFromPEM(cert)
	if !status {
		log.Error("certificate pool could not be updated")
		return err
	}
	conn, err := grpc.Dial(mc.Address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, "")))
	if err != nil {
		log.Error("could not send request to master")
		return err
	}
	defer conn.Close()
	if err := nodeconfig.AddNodeClient(conn); err != nil {
		log.Error("error from AddNodeClient", err)
		return err
	}

	return nil
}

func updateHostConfig() (*nodeconfig.HostConfigModel, error) {
	if !app.MainFlagVal.OverwriteHostConfig {
		hcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
		if err != nil {
			log.Error("error reading hostconfig, overwrite is not passed: ", err)
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
	if err := addNode(hcM.GRPCAuth.TLSCertFile); err != nil {
		log.Error("error adding node to cluster", err)
		//return nil, err
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
