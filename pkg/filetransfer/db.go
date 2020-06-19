package filetransfer

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/navinds25/styx/pkg/sftp"
)

// BadgerDB is the DB instance for BadgerDB
type BadgerDB struct {
	FilesDB        *badger.DB
	FileTransferDB *badger.DB
}

// DataStore is the struct containing the FilesDB and ConfigDB Interfaces.
type DataStore struct {
	Files        FilesStore
	FileTransfer FTStore
}

// Data is the instance of DataStore
var Data DataStore

// FilesStore is the main interface for the backend
type FilesStore interface {
	CheckFileExists([]byte) (bool, error)
	AddFile(string, *sftp.TransferConfig) error
	GetFile() error
	DeleteFile(string) error
	CloseFilesDB() error
}

// InitFilesDB initializes the NodeConfigDB
func InitFilesDB(s FilesStore) {
	Data.Files = s
}

// Close close the database.
func (badgerDB BadgerDB) Close() error {
	if err := badgerDB.FilesDB.Close(); err != nil {
		return err
	}
	return nil
}

// FTStore is the interface for the FileTransfer DB
type FTStore interface {
	CloseFTDB() error
}

// InitFTDB initializes the FileTransfer DB
func InitFTDB(s FTStore) {
	Data.FileTransfer = s
}

// CloseFTDB closes the FileTransfer Database
func (badgerDB BadgerDB) CloseFTDB() error {
	if err := badgerDB.FileTransferDB.Close(); err != nil {
		return err
	}
	return nil
}
