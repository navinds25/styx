package sftpdata

import (
	"github.com/dgraph-io/badger"
)

// BadgerDB is the DB instance for BadgerDB
type BadgerDB struct {
	ConfigDB *badger.DB
	FilesDB  *badger.DB
}

// DataStore is the struct containing the FilesDB and ConfigDB Interfaces.
type DataStore struct {
	Config ConfigStore
	Files  FilesStore
}

// Data is the instance of DataStore
var Data DataStore

// Close close the database.
func (badgerDB BadgerDB) Close() error {
	if err := badgerDB.ConfigDB.Close(); err != nil {
		return err
	}
	if err := badgerDB.FilesDB.Close(); err != nil {
		return err
	}
	return nil
}
