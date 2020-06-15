package nodeconfig

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/ghodss/yaml"
)

// GRPCAuth is sub config for GRPC Auth details
type GRPCAuth struct {
	TLSCertFile    string `json:"tls_cert_file,omitempty"`
	TLSCertData    string `json:"tls_cert_data,omitempty"`
	TLSCertBinData []byte `json:"tls_cert_bin_data,omitempty"`
	TLSKeyFile     string `json:"tls_key_file,omitempty"`
	TLSKeyData     string `json:"tls_key_data,omitempty"`
	TLSKeyBinData  []byte `json:"tls_key_bin_data,omitempty"`
	Token          string `json:"token,omitempty"`
}

// SFTPAuth is the sub config for SFTP Auth details
type SFTPAuth struct {
	SFTPAuthType   string `json:"sftp_auth_type,omitempty"` // SFTPAuthType is the type of authentication: password, key or key & passphrase
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	KeyFile        string `json:"key_file,omitempty"`
	KeyData        string `json:"key_data,omitempty"`
	KeyBinData     []byte `json:"key_bin_data,omitempty"`
	HostkeyFile    string `json:"hostkey_file,omitempty"`
	HostkeyData    string `json:"hostkey_data,omitempty"`
	HostkeyBinData []byte `json:"hostkey_bin_data,omitempty"`
}

// HostConfigInput holds the yaml configuration of the styxnode from user input/config file
// this includes data the other styxnodes do not require for peer to peer communication
type HostConfigInput struct {
	NodeID         string   `json:"node_id"`             // NodeID is unique identifier for a styxnode
	NodeType       string   `json:"node_type,omitempty"` // NodeType indicates if it's a internal/external node
	IPAddress      string   `json:"ip_address,omitempty"`
	GRPCPort       int      `json:"grpc_port,omitempty"`
	SFTPPort       int      `json:"sftp_port,omitempty"`
	SZ             string   `json:"sz,omitempty"` // SZ indicates the security zone, eg: dmz or control plane
	GRPCAuth       GRPCAuth `json:"grpc_auth,omitempty"`
	SFTPAuth       SFTPAuth `json:"sftp_auth,omitempty"`
	ExternalAccess bool     `json:"external_access,omitempty"` // ExternalAccess indicates the styxnode is allowed to send files outside, eg: external sftp server
	//Overwrite      bool `json:"overwrite"`
	EncryptionKey string            `json:"encryption_key"`
	MasterConfig  MasterConfigInput `json:"master_config,omitempty"`
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
func HostConfigToModel(hc *HostConfigInput) (*HostConfigModel, *MasterConfigModel, error) {
	//model := HostConfigModel(*hc)
	model := HostConfigModel{
		NodeID:    hc.NodeID,
		NodeType:  hc.NodeType,
		IPAddress: hc.IPAddress,
		GRPCPort:  hc.GRPCPort,
		SFTPPort:  hc.SFTPPort,
		SZ:        hc.SZ,
		GRPCAuth: &GRPCAuthModel{
			TLSCertFile: hc.GRPCAuth.TLSCertFile,
			TLSKeyFile:  hc.GRPCAuth.TLSKeyFile,
		},
		SFTPAuth: &SFTPAuthModel{
			SFTPAuthType: hc.SFTPAuth.SFTPAuthType,
			Username:     hc.SFTPAuth.Username, Password: hc.SFTPAuth.Password,
			KeyFile:     hc.SFTPAuth.KeyFile,
			KeyData:     hc.SFTPAuth.KeyData,
			HostkeyFile: hc.SFTPAuth.HostkeyFile,
		},
		ExternalAccess: hc.ExternalAccess,
		GRPCAddress:    hc.IPAddress + ":" + strconv.Itoa(hc.GRPCPort),
		SFTPAddress:    hc.IPAddress + ":" + strconv.Itoa(hc.SFTPPort),
	}
	masterConfig := MasterConfigModel{
		Master:  hc.MasterConfig.Master,
		Address: hc.MasterConfig.MasterIP + ":" + strconv.Itoa(hc.MasterConfig.MasterPort),
	}
	return &model, &masterConfig, nil
}
