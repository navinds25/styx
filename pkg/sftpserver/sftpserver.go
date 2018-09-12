package sftpserver

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

// GetConfig returns the config for the ssh/sftp server
func GetConfig() *ssh.ServerConfig {
	config := ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in
			// a production setting.
			log.Info("Login: %s\n", c.User())
			if c.User() == "testuser" && string(pass) == "tiger" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	privateBytes, err := ioutil.ReadFile("/home/navin/go/src/github.com/navinds25/styx/ssh_host_rsa_key")
	if err != nil {
		log.Fatal("Failed to load private key", err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key", err)
	}

	config.AddHostKey(private)

	return &config
}
