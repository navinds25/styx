package styxdata

import (
	"errors"

	"github.com/dgraph-io/badger"
	"github.com/navinds25/styx/pkg/styxsftp"
)

// InitConfigDB Initializes the Database
func InitConfigDB(s ConfigStore) {
	Data.Config = s
}

// ConfigStore is the main interface for the backend
type ConfigStore interface {
	CheckConfigExists([]byte) (bool, error)
	AddNodeEntry() error
	AddSFTPEntry(*styxsftp.TransferConfig) error
	DeleteSFTPEntry(string) error
	//UpdateSFTPEntry()
	GetAll() ([]styxsftp.TransferConfig, error)
	CloseConfigDB() error
}

// CloseConfigDB closes the database
// This is because we are not setting up the DB from the main function
func (badgerDB BadgerDB) CloseConfigDB() error {
	if err := badgerDB.ConfigDB.Close(); err != nil {
		return err
	}
	return nil
}

// CheckConfigExists checks for a key in the DB
func (badgerDB BadgerDB) CheckConfigExists(id []byte) (bool, error) {
	tx := badgerDB.ConfigDB.NewTransaction(false)
	defer tx.Discard()
	item, err := tx.Get(id)
	if err.Error() == "ErrKeyNotFound" {
		return false, nil
	} else if err != nil {
		return false, err
	} else if item != nil {
		return true, nil
	} else {
		return false, errors.New("Unknown error")
	}
}

// AddSFTPEntry is for adding a SFTP Config entry.
func (badgerDB BadgerDB) AddSFTPEntry(config *styxsftp.TransferConfig) error {
	//if err := config.EncryptSecureFields(); err != nil {
	//	return err
	//}
	value, err := config.EncodeGob()
	if err != nil {
		return err
	}
	txn := badgerDB.ConfigDB.NewTransaction(true)
	defer txn.Discard()
	if err := txn.Set([]byte(config.TransferID), value); err != nil {
		return err
	}
	err = badgerDB.ConfigDB.Update(func(txn *badger.Txn) error {
		if err := txn.Set([]byte(config.TransferID), value); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// DeleteSFTPEntry takes the TransferID and deletes the corresponding config.
func (badgerDB BadgerDB) DeleteSFTPEntry(id string) error {
	err := badgerDB.ConfigDB.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		return err
	})
	return err
}

// GetAll returns the entire configuration. TODO: use streams
func (badgerDB BadgerDB) GetAll() ([]styxsftp.TransferConfig, error) {
	id1 := styxsftp.TransferConfig{
		TransferID: "id1",
	}
	allConfig := []styxsftp.TransferConfig{}
	allConfig = append(allConfig, id1)
	return allConfig, nil
}
