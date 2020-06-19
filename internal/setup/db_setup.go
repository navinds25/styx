package setup

import (
	"encoding/base64"
	"os"
	"path/filepath"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/filetransfer"
	"github.com/navinds25/styx/pkg/nodeconfig"
	log "github.com/sirupsen/logrus"
)

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

func nodeConfigDB(dataRoot string) error {
	nodeConfigDBPath := filepath.Join(dataRoot, "nodeconfig")
	if err := createDataDir(nodeConfigDBPath); err != nil {
		return err
	}
	encryptionKey, err := base64.StdEncoding.DecodeString(app.MainFlagVal.EncryptionKey)
	if err != nil {
		log.Error("Invalid encryption key", err)
		return err
	}
	// WithTruncate is needed for windows
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

func filetransferDB(dataRoot string) error {
	dbPath := filepath.Join(dataRoot, "filetransfer")
	if err := createDataDir(dbPath); err != nil {
		return err
	}
	// WithTruncate is needed for windows
	dbOpts := badger.DefaultOptions(dbPath).WithTruncate(true).WithLogger(nil)
	dBInst, err := badger.Open(dbOpts)
	if err != nil {
		return err
	}
	filetransfer.InitFTDB(filetransfer.BadgerDB{
		FileTransferDB: dBInst,
	})
	return nil
}

// DBsetup sets up all the databases required
func DBsetup(dataRoot string) error {
	if err := nodeConfigDB(dataRoot); err != nil {
		return err
	}
	if err := filetransferDB(dataRoot); err != nil {
		return err
	}
	return nil
}
