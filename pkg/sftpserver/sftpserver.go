package sftpserver

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"

	"github.com/pkg/sftp"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

// Run starts a sftpserver
func Run() {
	config := &ssh.ServerConfig{
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

	ln, err := net.Listen("tcp", "0.0.0.0:2888")
	if err != nil {
		log.Error("Failed to listen", err)
	}
	defer ln.Close()
	nConn, err := ln.Accept()
	_, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		log.Error("Error creating server: ", err)
	}
	log.Info("SSH Server established")
	go ssh.DiscardRequests(reqs)

	for newChannel := range chans {
		log.Info("Incoming Channel: ", newChannel.ChannelType())
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Error("could not accept channel", err)
		}
		go func(in <-chan *ssh.Request) {
			for req := range in {
				ok := false
				switch req.Type {
				case "subsystem":
					log.Info("Subsystem: ", string(req.Payload[4:]))
					if string(req.Payload[4:]) == "sftp" {
						ok = true
					}
				}
				req.Reply(ok, nil)
			}
		}(requests)

		server, err := sftp.NewServer(channel, sftp.ReadOnly())
		if err != nil {
			log.Error(err)
		}

		//sftp.Handlers{}
		//server, err := sftp.NewRequestServer(channel)

		if err := server.Serve(); err == io.EOF {
			log.Info("sftp client exited session.")
		} else if err != nil {
			log.Error("sftp server completed with error", err)
		}

		defer server.Close()
	}

}
