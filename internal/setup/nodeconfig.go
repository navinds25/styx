package setup

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/nodeconfig"
	log "github.com/sirupsen/logrus"
)

func nodeConfigDB(dataRoot string) error {
	nodeConfigDBPath := filepath.Join(dataRoot, "node_config")
	if err := createDataDir(nodeConfigDBPath); err != nil {
		return err
	}
	base64.StdEncoding.DecodeString(app.MainFlagVal.EncryptionKey)
	dbOpts := badger.DefaultOptions(nodeConfigDBPath)
	dbOpts.WithEncryptionKey([]byte(""))
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

func overwriteFromCli() (*nodeconfig.HostConfigModel, error) {
	cliHC, err := readHostConfigFromCli()
	if cliHC == nil || err != nil {
		if cliHC == nil {
			return nil, errors.New("cli didn't have any value")
		}
		return nil, err
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
		return hcM, err
	}
	log.Debug("got entry from DB", dHcM)
	return dHcM, nil
}

func updateHostConfig() (*nodeconfig.HostConfigModel, error) {
	if !app.MainFlagVal.OverwriteHostConfig {
		hcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
		if err != nil {
			log.Error("error reading hostconfig")
			return nil, err
		}
		log.Debugf("got hostconfig from db: %v", hcM)
		return hcM, nil
	}
	hcM, err := overwriteFromCli()
	if err != nil {
		log.Error("error updating hostconfig", err)
	}
	log.Debugf("got hostconfig from db/overwrite: %v", hcM)
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
		log.Error("error updating hostconfig", err)
		return nil, err
	}
	return hcM, nil
}
