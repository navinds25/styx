package sftpconfig

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/navinds25/styx/pkg/encryption"
)

// RawEncryptionKey is for the AES encryption key, it is base64Encoded.
var RawEncryptionKey = "SlNqVVhZalpzWWg5ZkRWOENvWHdwNzhITnl3RnpnWnFE"

// OpEnum is the Enum for encryption operations
var OpEnum = &Op{
	Encrypt: 1,
	Decrypt: 0,
}

// Op is for encryption operation
type Op struct {
	Encrypt int
	Decrypt int
}

// Config is the Interface for methods on TransferConfig struct
type Config interface {
	EncryptSecureFields(key string) error
	DecryptSecureFields(key string) error
	EncodeGob() error
	DecodeGob() error
}

// TransferConfig is the struct for parsing the configs
type TransferConfig struct {
	TransferID        string `json:"transfer_id"`
	Description       string `json:"description"`
	Type              string `json:"type"`
	LocalFile         string `json:"local_file"`
	LocalPath         string `json:"local_path"`
	RemoteFile        string `json:"remote_file"`
	RemotePath        string `json:"remote_path"`
	RemoteHost        string `json:"remote_host"`
	RemotePort        int    `json:"remote_port"`
	RemoteUser        string `json:"remote_user"`
	RemotePassword    string `json:"remote_password"`
	RemoteAuthKeyFile string `json:"remote_auth_key"`
}

func aesString(value string, op int) (string, error) {
	key, err := base64.StdEncoding.DecodeString(RawEncryptionKey)
	if err != nil {
		return "", err
	}
	if op == 1 {
		ciphertext, err := encryption.AESEncryptCBC(key, []byte(value))
		if err != nil {
			return "", err
		}
		return string(ciphertext), nil
	} else if op == 0 {
		ciphertext, err := encryption.AESDecryptCBC(key, value)
		if err != nil {
			return "", err
		}
		return string(ciphertext), nil
	}
	return "", nil
}

// EncryptSecureFields encrypts the fields of struct instance.
func (config *TransferConfig) EncryptSecureFields() error {
	op := OpEnum.Encrypt
	if err := secureFields(config, op); err != nil {
		return err
	}
	return nil
}

// DecryptSecureFields decrypts the fields of the struct instance.
func (config *TransferConfig) DecryptSecureFields() error {
	op := OpEnum.Decrypt
	if err := secureFields(config, op); err != nil {
		return err
	}
	return nil
}

func secureFields(config *TransferConfig, op int) error {
	var err error
	user := config.RemoteUser
	config.RemoteUser, err = aesString(user, op)
	if err != nil {
		return err
	}
	password := config.RemotePassword
	config.RemotePassword, err = aesString(password, op)
	if err != nil {
		return err
	}
	return nil
}

// EncodeGob Returns Gob encoded byte array of struct
func (config *TransferConfig) EncodeGob() ([]byte, error) {
	value := bytes.Buffer{}
	if err := gob.NewEncoder(&value).Encode(config); err != nil {
		return nil, err
	}
	return value.Bytes(), nil
}

// DecodeGob takes a encoded Gob as a byte array and updates the
// struct instance with the decoded values
func (config *TransferConfig) DecodeGob(value []byte) error {
	valReader := bytes.NewReader(value)
	if err := gob.NewDecoder(valReader).Decode(config); err != nil {
		return err
	}
	return nil
}

// GetConfig returns the parsed config
func GetConfig(confFile string) (map[string][]TransferConfig, error) {
	var config map[string][]TransferConfig
	_, err := os.Stat(confFile)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, err
	}
	return config, nil
}
