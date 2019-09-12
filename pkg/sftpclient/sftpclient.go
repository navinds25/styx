package sftpclient

import (
	"errors"
	"io/ioutil"

	cbcssh "github.com/ScriptRock/crypto/ssh"
	cbcsftp "github.com/ScriptRock/sftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client is a struct for client object + properties
type Client struct {
	Conn    *sftp.Client
	CBCConn *cbcsftp.Client
	CBC     bool
}

// Input is a struct for creating the sftp client
type Input struct {
	Address  string
	Username string
	Password string
	Protocol string
	// AuthMethod: One of pk, pk+pass, pass
	AuthMethod string
	PrivateKey string
	// ?
	ConnectionType string
	CBC            bool
}

// Regular Library:

// CreateSSHConfig returns a config instance used to create an ssh/sftp connection.
func CreateSSHConfig(inputConf *Input) (*ssh.ClientConfig, error) {
	switch inputConf.AuthMethod {
	case "pk":
		key, err := ioutil.ReadFile(inputConf.PrivateKey)
		if err != nil {
			return nil, err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		config := &ssh.ClientConfig{
			User: inputConf.Username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		return config, nil
	case "pk+pass":
		return nil, errors.New("Not Implemented")

	default:
		config := &ssh.ClientConfig{
			User: inputConf.Username,
			Auth: []ssh.AuthMethod{
				ssh.Password(inputConf.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		return config, nil
	}
}

// CBC Library

// CBCCreateSSHConfig returns a config instance used to create an ssh/sftp connection.
func CBCCreateSSHConfig(inputConf *Input) (*cbcssh.ClientConfig, error) {
	switch inputConf.AuthMethod {
	case "pk":
		key, err := ioutil.ReadFile(inputConf.PrivateKey)
		if err != nil {
			return nil, err
		}
		signer, err := cbcssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		config := &cbcssh.ClientConfig{
			User: inputConf.Username,
			Auth: []cbcssh.AuthMethod{
				cbcssh.PublicKeys(signer),
			},
			HostKeyCallback: cbcssh.InsecureIgnoreHostKey(),
		}
		return config, nil
	case "pk+pass":
		return nil, errors.New("Not Implemented")

	default:
		config := &cbcssh.ClientConfig{
			User: inputConf.Username,
			Auth: []cbcssh.AuthMethod{
				cbcssh.Password(inputConf.Password),
			},
			HostKeyCallback: cbcssh.InsecureIgnoreHostKey(),
		}
		return config, nil
	}
}

// CreateClient creates an sftp client
func CreateClient(inputConf *Input) (*Client, error) {
	client := Client{}
	if inputConf.CBC {
		config, err := CBCCreateSSHConfig(inputConf)
		if err != nil {
			return nil, err
		}
		config.Config.Ciphers = append(config.Config.Ciphers, "aes256-cbc")
		conn, err := cbcssh.Dial(inputConf.Protocol, inputConf.Address, config)
		if err != nil {
			return nil, err
		}
		client.CBCConn, err = cbcsftp.NewClient(conn)
		if err != nil {
			return nil, err
		}
		client.CBC = inputConf.CBC
		return &client, nil
	}
	config, err := CreateSSHConfig(inputConf)
	if err != nil {
		return nil, err
	}
	conn, err := ssh.Dial(inputConf.Protocol, inputConf.Address, config)
	if err != nil {
		return nil, err
	}
	client.Conn, err = sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	client.CBC = inputConf.CBC
	return &client, nil
}
