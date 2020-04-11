package nodeconfig

import "github.com/dgraph-io/badger/v2"

// BadgerDB is the DB instance for BadgerDB
type BadgerDB struct {
	NodeConfigDB *badger.DB
}

// DataStore is the struct containing the NodeConfigStore interface
type DataStore struct {
	NodeConfig Store
}

// Data is the instance of DataStore
var Data DataStore

// Store is the interface for all NodeConfig DB Actions
type Store interface {
	Close() error
}

// InitNodeConfigDB initializes the NodeConfigDB
func InitNodeConfigDB(s Store) {
	Data.NodeConfig = s
}

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
