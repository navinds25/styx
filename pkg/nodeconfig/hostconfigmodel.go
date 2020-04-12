package nodeconfig

import (
	"bytes"
	"encoding/gob"
	"strings"
)

// HostConfigKey is the key name for the hostconfig in the DB
const HostConfigKey = "hostconfig"

// HostConfigModel holds the configuration of the styxnode as stored in the db
// this includes data the other styxnodes do not require for peer to peer communication
type HostConfigModel struct {
	NodeID    string `json:"node_id"`   // NodeID is unique identifier for a styxnode
	NodeType  string `json:"node_type"` // NodeType indicates if it's a internal/external node
	IPAddress string `json:"ipaddress"`
	GrpcPort  string `json:"grpc_port"`
	SftpPort  string `json:"sftp_port"`
	EnvSec    string `json:"envsec"` // EnvSec indicates the environment tier, eg: dmz or control plane
	GrpcAuth  struct {
		TLSCertFile    string `json:"tls_cert_file"`
		TLSCertData    string `json:"tls_cert_data"`
		TLSCertBinData []byte
		TLSKeyFile     string `json:"tls_key_file"`
		TLSKeyData     string `json:"tls_key_data"`
		TLSKeyBinData  []byte
	} `json:"grpc_auth"`
	SftpAuth struct {
		SFTPAuthType   string `json:"sftp_auth_type"` // SFTPAuthType is the type of authentication: password, key or key & passphrase
		Username       string `json:"username"`
		Password       string `json:"password"`
		KeyFile        string `json:"key_file"`
		KeyData        string `json:"key_data"`
		KeyBinData     []byte
		HostkeyFile    string `json:"hostkey_file"`
		HostkeyData    string `json:"hostkey_data"`
		HostkeyBinData []byte
	} `json:"sftp_auth"`
	ExternalAccess bool `json:"external_access"` // ExternalAccess indicates the styxnode is allowed to send files outside, eg: external sftp server
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
func (badgerDB BadgerDB) GetHostConfigEntry(id string) (outModel *HostConfigModel, exists bool, err error) {
	txn := badgerDB.NodeConfigDB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte(id))
	if err != nil {
		if err.Error() != "ErrKeyNotFound" {
			exists = true
		}
		return nil, exists, err
	}
	exists = true
	tmpVal := []byte{}
	tmpVal, err = item.ValueCopy(tmpVal)
	if err != nil {
		return nil, exists, err
	}
	if err := gob.NewDecoder(bytes.NewReader(tmpVal)).Decode(outModel); err != nil {
		return nil, exists, err
	}
	return outModel, exists, nil
}
