package main

import (
	"net"
	"os"

	"github.com/navinds25/styx/internal/grpcdef"
	"github.com/navinds25/styx/internal/styxcli"
	"github.com/navinds25/styx/pkg/sftpserver"
	"github.com/navinds25/styx/pkg/styxconfig"
	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func message() {
	log.Info("Closing stuff")
}

func shutdown() {
	log.Info("Exiting cleanly")
	os.Exit(0)
}

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

func main() {
	app := styxcli.App()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	config, err := styxconfig.GetConfig("_extra/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)
	//t := tebata.New(syscall.SIGINT, syscall.SIGTERM)
	grpcAddress := styxcli.InterfaceAddress + ":" + styxcli.GrpcPort
	grpclis, err := net.Listen("tcp", grpcAddress)
	defer grpclis.Close()
	sftpAddress := styxcli.InterfaceAddress + ":" + styxcli.SftpPort
	sftplis, err := net.Listen("tcp", sftpAddress)
	defer sftplis.Close()
	if err != nil {
		log.Fatal(err)
	}
	//t.Reserve(sftplis.Close)
	// for grpc server
	log.Infof("Listening grpc server %s ,sftp server: %s",
		grpcAddress,
		sftpAddress,
	)
	s := grpc.NewServer()
	//t.Reserve(s.GracefulStop)
	go func() {
		if err := runGRPCServer(grpclis, s); err != nil {
			log.Fatal(err)
		}
	}()
	//t.Reserve(sftplis.Close)
	//t.Reserve(shutdown)

	// for sftp server
	//if err := listenSFTPServer(t, sftplis); err != nil {
	if err := sftpserver.ListenSFTPServer(sftplis); err != nil {
		log.Fatal("Error from main", err)
	}
}
