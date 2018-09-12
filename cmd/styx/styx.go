package main

import (
	"io"
	"net"
	"os"
	"syscall"

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
	grpcport = ":50051"
	sftpport = ":28888"
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

func runSFTPServer(t *tebata.Tebata, lis net.Listener) error {
	//t.Reserve(message)
	nConn, err := lis.Accept()
	config := sftpserver.GetConfig()
	_, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		return err
	}
	go ssh.DiscardRequests(reqs)

	for newChannel := range chans {
		log.Info("Incoming Channel: ", newChannel.ChannelType())
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			return err
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
			return err
		}
		//sftp.Handlers{}
		//server, err := sftp.NewRequestServer(channel)
		defer server.Close()
		// t.Reserve(server.Close) - doesn't work
		if err := server.Serve(); err == io.EOF {
			log.Info("sftp client exited session.")
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
	if err := runSFTPServer(t, sftplis); err != nil {
		log.Fatal(err)
	}
}
