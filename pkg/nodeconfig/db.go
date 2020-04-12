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
	AddHostConfigEntry(string, *HostConfigModel) error
	GetHostConfigEntry(string) (*HostConfigModel, bool, error)
	CloseDB() error
}

// InitNodeConfigDB initializes the NodeConfigDB
func InitNodeConfigDB(s Store) {
	Data.NodeConfig = s
}

// CloseDB close the database.
func (badgerDB BadgerDB) CloseDB() error {
	if err := badgerDB.NodeConfigDB.Close(); err != nil {
		return err
	}
	return nil
}
