package app

import (
	"os"
	"path"

	"github.com/dgraph-io/badger"
	"github.com/navinds25/styx/internal/styxcli"
	"github.com/navinds25/styx/pkg/sftpdata"
	log "github.com/sirupsen/logrus"
)

func createDataDir(datadir string) error {
	log.Debug("datadir:", datadir)
	if err := os.MkdirAll(datadir, 0755); err != nil {
		return err
	}
	log.Info("created data directory: ", datadir)
	return nil
}

// StyxNodeDBSetup sets up and initializes the DBs for the Styxnode
func StyxNodeDBSetup() error {
	return nil
}

// StyxMasterDBSetup sets up and initializes the DBs for Styxmaster.
func StyxMasterDBSetup() error {
	return nil
}

// SFTPDBSetup sets up and initializes the DBs for External SFTP Management.
func SFTPDBSetup() error {
	log.Info("Starting DB Setup")
	log.Debug("data directory is: ", styxcli.MainFlagVal.DataDir)
	configDBDir := path.Join(styxcli.MainFlagVal.DataDir, "config")
	filesDBDir := path.Join(styxcli.MainFlagVal.DataDir, "files")
	_, Err := os.Stat(styxcli.MainFlagVal.DataDir)
	if Err != nil {
		if err := createDataDir(styxcli.MainFlagVal.DataDir); err != nil {
			return err
		}
	}
	configDBopts := badger.DefaultOptions
	configDBopts.Dir = configDBDir
	configDBopts.ValueDir = path.Join(configDBDir, "value")
	//configDBopts.Logger = log.StandardLogger()
	configDB, err := badger.Open(configDBopts)
	if err != nil {
		return err
	}
	sftpdata.InitConfigDB(sftpdata.BadgerDB{
		ConfigDB: configDB,
	})
	filesDBopts := badger.DefaultOptions
	filesDBopts.Dir = filesDBDir
	filesDBopts.ValueDir = path.Join(filesDBDir, "value")
	//filesDBopts.Logger = log.StandardLogger()
	filesDB, err := badger.Open(filesDBopts)
	if err != nil {
		return err
	}
	sftpdata.InitFilesDB(sftpdata.BadgerDB{
		FilesDB: filesDB,
	})
	log.Info("Databases initialized.")
	return nil
}
