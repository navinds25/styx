package sftpclient

import (
	"errors"
	"io/ioutil"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client is a struct for client object + properties
type Client struct {
	Conn *sftp.Client
	ID   string
}

// Input is a struct for creating the sftp client
type Input struct {
	Address        string
	Username       string
	Password       string
	Protocol       string
	AuthMethod     string
	PrivateKey     string
	ConnectionType string
}

// CreateSSHConfig returns a config instance used to create an ssh/sftp connection.
func CreateSSHConfig(inputConf Input) (*ssh.ClientConfig, error) {
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

// CreateClient creates an sftp client
func CreateClient(inputConf Input) (*Client, error) {
	client := Client{}
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

	/*
		_, err = client.SFTPConn.Getwd()
		if err != nil {
			return nil, err
		}

	*/
	return &client, nil
}
