package main

import (
	"io"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/pkg/sftp"
	"github.com/syossan27/tebata"
	"golang.org/x/crypto/ssh"

	"github.com/navinds25/styx/pkg/grpcdef"
	"github.com/navinds25/styx/pkg/sftpserver"
	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	grpcport = "127.0.0.1:50051"
	sftpport = "127.0.0.1:28888"
)

func runGRPCServer(lis net.Listener, s *grpc.Server) error {
	pb.RegisterStyxServer(s, &grpcdef.Server{})
	log.Info("Started GRPC Server")
	if err := s.Serve(lis); err != nil {
		return err
	}
	defer s.GracefulStop()
	defer lis.Close()
	return nil
}

func listenSFTPServer(t *tebata.Tebata, lis net.Listener) error {
	//t.Reserve(message)
	for {
		t.Reserve(lis.Close)
		nConn, err := lis.Accept()
		if err != nil {
			log.Error("Error from listenSFTPServer:", err)
			return (err)
		}
		go runSFTPServer(nConn)
	}
}
func runSFTPServer(nConn net.Conn) error {
	nConn.SetReadDeadline(time.Now().Add(5 * time.Second))
	nConn.SetDeadline(time.Now().Add(600 * time.Second))
	nConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	config := sftpserver.GetConfig()
	// set timeout incase client opens connection but doesn't do anything.
	serverConn, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		return err
	}
	go ssh.DiscardRequests(reqs)

	for newChannel := range chans {
		log.Info("Incoming Channel: ", newChannel.ChannelType())
		if newChannel.ChannelType() != "session" {
			log.Info("Unknown: ", newChannel.ChannelType())
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		} else {
			log.Info("Got channel: ", newChannel.ChannelType())
			log.Info("Extra info:", string(newChannel.ExtraData()))
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			return err
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
					ok = false
				}
				req.Reply(ok, nil)
			}
		}(requests)
		serverOptions := []sftp.ServerOption{
			sftp.WithDebug(os.Stdout),
			sftp.ReadOnly(),
		}

		server, err := sftp.NewServer(channel, serverOptions...)
		if err != nil {
			return err
		}
		//sftp.Handlers{}
		//server, err := sftp.NewRequestServer(channel)
		defer server.Close()
		// t.Reserve(server.Close) - doesn't work
		if err := server.Serve(); err == io.EOF {
			log.Infof("sftp client %s exited session.", serverConn.ClientVersion())
			server.Close()
		} else if err != nil {
			log.Error("sftp server completed with error", err)
		}
	}
	return nil
}

func message() {
	log.Info("Closing stuff")
}

func shutdown() {
	log.Info("Exiting cleanly")
	os.Exit(0)
}

func main() {
	t := tebata.New(syscall.SIGINT, syscall.SIGTERM)
	grpclis, err := net.Listen("tcp", grpcport)
	defer grpclis.Close()
	sftplis, err := net.Listen("tcp", sftpport)
	defer sftplis.Close()
	if err != nil {
		log.Fatal(err)
	}
	//t.Reserve(sftplis.Close)
	// for grpc server
	log.Infof("Listening for grpc on: %s ; sftp on %s", grpcport, sftpport)
	s := grpc.NewServer()
	t.Reserve(s.GracefulStop)
	go func() {
		if err := runGRPCServer(grpclis, s); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()
	t.Reserve(sftplis.Close)
	t.Reserve(shutdown)

	// for sftp server
	if err := listenSFTPServer(t, sftplis); err != nil {
		log.Error("Error from main", err)
	}
}
