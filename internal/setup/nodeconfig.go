package setup

import (
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v2"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/nodeconfig"
	log "github.com/sirupsen/logrus"
)

type ListenConfig struct {
	GRPCAddress string
	SFTPAddress string
}

func nodeConfigDB(dataRoot string) error {
	nodeConfigDBPath := filepath.Join(dataRoot, "node_config")
	if err := createDataDir(nodeConfigDBPath); err != nil {
		return err
	}
	dbOpts := badger.DefaultOptions(nodeConfigDBPath)
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

func checkHostConfigExists() (*nodeconfig.HostConfigModel, error) {
	// 1. check db for host config 2. check for overwrite flag
	hcM, exists, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
	if err != nil {
		return nil, err
	}
	// if 1 is absent and 2 is present -> write to db & send request to styxmaster
	if !exists || app.MainFlagVal.OverwriteHostConfig {
		if cliHC, err := readHostConfigFromCli(); cliHC != nil && err == nil {
			hcM, err := nodeconfig.HostConfigToModel(cliHC)
			if err != nil {
				return nil, err
			}
			if err := nodeconfig.Data.NodeConfig.AddHostConfigEntry(nodeconfig.HostConfigKey, hcM); err != nil {
				return nil, err
			}
			hcM, _, err = nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
			if err != nil {
				return hcM, err
			}
		}
	}
	return hcM, nil
}

// NodeSetup is the main setup function
func NodeSetup() (hcM *nodeconfig.HostConfigModel, lis *ListenConfig, err error) {
	// read cli flags
	app.MainFlagVal.CliSetDefaults()

	// setup dbs
	if err := nodeConfigDB(app.MainFlagVal.DataDir); err != nil {
		return hcM, lis, err
	}
	// get the hostconfig
	hcM, err = checkHostConfigExists()
	if err != nil {
		return hcM, lis, err
	}
	lis.GRPCAddress = hcM.IPAddress + ":" + hcM.GrpcPort
	lis.SFTPAddress = hcM.IPAddress + ":" + hcM.SftpPort
	return hcM, lis, err
}
