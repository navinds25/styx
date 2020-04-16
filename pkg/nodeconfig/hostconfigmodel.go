package nodeconfig

import (
	"bytes"
	"encoding/gob"
	"strings"
)

// HostConfigKey is the key name for the hostconfig in the DB
const HostConfigKey = "hostconfig"

// GRPCAuthModel is sub config for GRPC Auth details
type GRPCAuthModel struct {
	TLSCertFile    string
	TLSCertData    string
	TLSCertBinData []byte
	TLSKeyFile     string
	TLSKeyData     string
	TLSKeyBinData  []byte
}

// SFTPAuthModel is the sub config for SFTP Auth details
type SFTPAuthModel struct {
	SFTPAuthType   string // SFTPAuthType is the type of authentication: password, key or key & passphrase
	Username       string
	Password       string
	KeyFile        string
	KeyData        string
	KeyBinData     []byte
	HostkeyFile    string
	HostkeyData    string
	HostkeyBinData []byte
}

// HostConfigModel holds the configuration of the styxnode as stored in the db
// this includes data the other styxnodes do not require for peer to peer communication
type HostConfigModel struct {
	NodeID         string // NodeID is unique identifier for a styxnode
	NodeType       string // NodeType indicates if it's a internal/external node
	IPAddress      string
	GRPCPort       int
	SFTPPort       int
	SZ             string // SZ indicates the security zone, eg: dmz or control plane
	GRPCAuth       GRPCAuthModel
	SFTPAuth       SFTPAuthModel
	ExternalAccess bool // ExternalAccess indicates the styxnode is allowed to send files outside, eg: external sftp server
	GRPCAddress    string
	SFTPAddress    string
}

// AddHostConfigEntry adds the HostConfigModel to DB
func (badgerDB BadgerDB) AddHostConfigEntry(id string, inModel *HostConfigModel) error {
	key := strings.TrimSpace(id)
	buf := bytes.Buffer{}
	if err := gob.NewEncoder(&buf).Encode(inModel); err != nil {
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

// GetHostConfigEntry gets HostConfigModel by id
// Returns entry, exists bool and error
func (badgerDB BadgerDB) GetHostConfigEntry(id string) (*HostConfigModel, error) {
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
	outModel := &HostConfigModel{} // gob doesn't let you decode to nil pointer
	if err := gob.NewDecoder(bytes.NewReader(tmpVal)).Decode(outModel); err != nil {
		return nil, err
	}
	return outModel, nil
}
