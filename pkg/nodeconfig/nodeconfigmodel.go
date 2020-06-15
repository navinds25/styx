package nodeconfig

import (
	"bytes"
	"encoding/gob"
	"strings"

	badger "github.com/dgraph-io/badger/v2"
	pb "github.com/navinds25/styx/api/nodeconfig"
)

// The model for nodeconfig is present in api/nodeconfig

// NodeConfigPrefixKey is the prefix for storing all the nodeconfigs in db
const NodeConfigPrefixKey = "nodeconfig"

// DBSeparator is pipe
const DBSeparator = "|"

// AddNodeConfigEntry adds a single nodeconfig entry to badgerDB
func (badgerDB BadgerDB) AddNodeConfigEntry(id string, entry *pb.NodeConfig) error {
	key := strings.TrimSpace(id)
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

// GetNodeConfigEntry returns a nodeconfig entry for the key
func (badgerDB BadgerDB) GetNodeConfigEntry(id string) (*pb.NodeConfig, error) {
	txn := badgerDB.NodeConfigDB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte(id))
	if err != nil {
		return nil, err
	}
	tmpVal, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	outModel := &pb.NodeConfig{} // gob doesn't let you decode to nil pointer
	if err := gob.NewDecoder(bytes.NewReader(tmpVal)).Decode(outModel); err != nil {
		return nil, err
	}
	return outModel, nil
}

// GetAllNodeConfigEntries returns all the NodeConfig entrys for the prefix
func (badgerDB BadgerDB) GetAllNodeConfigEntries(prefix string) ([]*pb.NodeConfig, error) {
	allNodeConfigs := []*pb.NodeConfig{}
	txn := badgerDB.NodeConfigDB.NewTransaction(false)
	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()
	keyPrefix := []byte(prefix)
	for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
		item := it.Item()
		tmpVal, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}
		outNodeConfig := &pb.NodeConfig{}
		if err := gob.NewDecoder(bytes.NewReader(tmpVal)).Decode(outNodeConfig); err != nil {
			return nil, err
		}
		allNodeConfigs = append(allNodeConfigs, outNodeConfig)
	}
	return allNodeConfigs, nil
}
