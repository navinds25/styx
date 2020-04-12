package nodeconfig

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	pb "github.com/navinds25/styx/api/nodeconfig"

	"github.com/ghodss/yaml"
)

// HostConfigInput holds the yaml configuration of the styxnode from user input/config file
// this includes data the other styxnodes do not require for peer to peer communication
type HostConfigInput struct {
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
	//Overwrite      bool `json:"overwrite"`
}

// HostConfigFromYAML returns a HostConfig from yaml
func HostConfigFromYAML(inputYaml io.Reader) (*HostConfigInput, error) {
	config := &HostConfigInput{}
	data, err := ioutil.ReadAll(inputYaml)
	if err != nil {
		return nil, err
	}
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, config); err != nil {
		return nil, err
	}
	return config, nil
}

// HostConfigFromYAMLFile returns a HostConfigInput from a Yaml File
func HostConfigFromYAMLFile(filename string) (*HostConfigInput, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	fileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return HostConfigFromYAML(fileReader)
}

// HostConfigToModel converts HostConfigInput to the Model
func HostConfigToModel(hc *HostConfigInput) (*HostConfigModel, error) {
	model := HostConfigModel(*hc)
	return &model, nil
}

// HostConfigToNodeConfig takes a HostConfig and returns a NodeConfig
func HostConfigToNodeConfig(*HostConfigInput) (*pb.NodeConfig, error) {
	return &pb.NodeConfig{}, nil
}
