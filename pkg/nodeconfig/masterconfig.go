package nodeconfig

import (
	"bytes"
	"encoding/gob"
)

// MasterConfigInput is the
type MasterConfigInput struct {
	Master     bool   `json:master,omitempty`
	MasterIP   string `json:"master_ip,omitempty"`
	MasterPort int    `json:"master_port,omitempty"`
}

// MasterConfigModel holds the configuration of the styx master
type MasterConfigModel struct {
	Master  bool
	Address string
}

// AddMasterConfigEntry adds a single nodeconfig entry to badgerDB
func (badgerDB BadgerDB) AddMasterConfigEntry(entry *MasterConfigModel) error {
	key := "master"
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(entry); err != nil {
		return err
	}
	txn := badgerDB.NodeConfigDB.NewTransaction(true)
	defer txn.Discard()
	if err := txn.Set([]byte(key), buf.Bytes()); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

// GetMasterConfigEntry returns a masterconfig entry
func (badgerDB BadgerDB) GetMasterConfigEntry() (*MasterConfigModel, error) {
	txn := badgerDB.NodeConfigDB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte("master"))
	if err != nil {
		return nil, err
	}
	tmpVal, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	outModel := &MasterConfigModel{} // gob doesn't let you decode to nil pointer
	if err := gob.NewDecoder(bytes.NewReader(tmpVal)).Decode(outModel); err != nil {
		return nil, err
	}
	return outModel, nil
}
