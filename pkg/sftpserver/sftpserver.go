package sftpserver

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/ScriptRock/sftp"
	log "github.com/sirupsen/logrus"
	"github.com/ScriptRock/crypto/ssh"
)

// GetConfig returns the config for the ssh/sftp server
func GetConfig(sshhostkey string) *ssh.ServerConfig {
	config := ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in
			// a production setting.
			log.Infof("Login: %s\n", c.User())
			if c.User() == "testusr" && string(pass) == "tiger" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	privateBytes, err := ioutil.ReadFile(sshhostkey)
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

// ListenSFTPServer starts an SFTP Server
func ListenSFTPServer(lis net.Listener, sshhostkey string) error {
	config := GetConfig(sshhostkey)
	for {
		nConn, err := lis.Accept()
		if err != nil {
			log.Error("Error from listenSFTPServer:", err)
			return err
		}
		defer nConn.Close()
		_, chans, reqs, err := ssh.NewServerConn(nConn, config)
		if err != nil {
			log.Error(err)
		}
		go ssh.DiscardRequests(reqs)
		go handleChannels(chans)
	}
}

func handleChannels(chans <-chan ssh.NewChannel) error {
	for newChannel := range chans {
		go handleChannel(newChannel)
	}
	return nil
}

func handleChannel(newChannel ssh.NewChannel) {
	log.Info("Incoming Channel: ", newChannel.ChannelType())
	if newChannel.ChannelType() != "session" {
		log.Info("Unknown: ", newChannel.ChannelType())
		newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
	} else {
		log.Info("Got channel: ", newChannel.ChannelType())
		log.Info("Extra info:", string(newChannel.ExtraData()))
	}
	channel, requests, err := newChannel.Accept()
	if err != nil {
		log.Error(err)
		return
	}
	go func(in <-chan *ssh.Request) {
		for req := range in {
			ok := false
			log.Printf("payload: %v", string(req.Payload))
			log.Printf("type: %v", req.Type)
			switch req.Type {
			case "subsystem":
				log.Info("Subsystem: ", string(req.Payload[4:]))
				if string(req.Payload[4:]) == "sftp" {
					ok = true
				}
			default:
				log.Error("no subsystem")
				ok = false
				return
			}
			req.Reply(ok, nil)
			return
		}
	}(requests)
	serverOptions := []sftp.ServerOption{
		sftp.WithDebug(os.Stdout),
		sftp.ReadOnly(),
	}
	server, err := sftp.NewServer(channel, serverOptions...)
	if err != nil {
		log.Error(err)
		return
	}
	//sftp.Handlers{}
	//server, err := sftp.NewRequestServer(channel)
	defer server.Close()
	// t.Reserve(server.Close) - doesn't work
	if err := server.Serve(); err == io.EOF {
		server.Close()
	} else if err != nil {
		log.Error("sftp server completed with error", err)
	}
}
