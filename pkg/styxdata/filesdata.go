package styxdata

import (
	"bytes"
	"encoding/gob"
	"os"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/navinds25/styx/pkg/styxsftp"

	log "github.com/sirupsen/logrus"
)

// InitFilesDB Initializes the Database
func InitFilesDB(s FilesStore) {
	Data.Files = s
}

// FilesStore is the main interface for the backend
type FilesStore interface {
	CheckFileExists([]byte) (bool, error)
	AddFile(string, *styxsftp.TransferConfig) error
	GetFile() error
	DeleteFile(string) error
	CloseFilesDB() error
}

// CloseFilesDB closes the database.
// This is because we are not setting up the DB from the main function.
func (badgerDB BadgerDB) CloseFilesDB() error {
	if err := badgerDB.FilesDB.Close(); err != nil {
		return err
	}
	return nil
}

// CheckFileExists checks if a file exists in the database.
func (badgerDB BadgerDB) CheckFileExists(key []byte) (bool, error) {
	txn := badgerDB.FilesDB.NewTransaction(false)
	defer txn.Discard()
	_, err := txn.Get(key)
	if err != nil {
		if err.Error() == "ErrKeyNotFound" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AddFile adds a new file in the Files DB
// Key is the full path of the destination file.
// Value is TransferConfig for the file.
func (badgerDB BadgerDB) AddFile(key string, value *styxsftp.TransferConfig) error {
	fileKey := strings.TrimSpace(key)
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		return err
	}
	txn := badgerDB.FilesDB.NewTransaction(true)
	defer txn.Commit()
	log.Debug("AddFile: key {string}: ", fileKey)
	if err := txn.Set([]byte(fileKey), buf.Bytes()); err != nil {
		return err
	}
	return nil
}

// GetFile gets a file from Files DB
func (badgerDB BadgerDB) GetFile() error {
	return nil
}

// DeleteFile removes a file from the Files DB
func (badgerDB BadgerDB) DeleteFile(key string) error {
	fileKey := strings.TrimSpace(key)
	txn := badgerDB.FilesDB.NewTransaction(true)
	defer txn.Commit()
	log.Debug("DeleteFile: key {string}: ", fileKey)
	if err := txn.Delete([]byte(fileKey)); err != nil {
		return err
	}
	return nil
}

// ListFiles lists all the files in the DB.
func (badgerDB BadgerDB) ListFiles() error {
	log.Info("Listing Files:")
	err := badgerDB.FilesDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			log.Println("key=", string(k))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func createDataDir(datadir string) error {
	log.Debug("datadir:", datadir)
	if err := os.MkdirAll(datadir, 0755); err != nil {
		return err
	}
	log.Debug("created data directory: ", datadir)
	return nil
}
