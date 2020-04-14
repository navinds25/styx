package main

import (
	"net"
	"os"

	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/internal/setup"
	"github.com/navinds25/styx/pkg/nodeconfig"
	"github.com/navinds25/styx/pkg/sftp"
	log "github.com/sirupsen/logrus"
)

// Version for inserting version via ldflags
var Version string

func main() {
	setup.Logging()
	appCli := app.Cli()
	appCli.Version = Version
	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	_, lis, err := setup.NodeSetup()
	if err != nil {
		log.Fatal(err)
	}
	defer nodeconfig.Data.NodeConfig.CloseDB()
	grpcListener, err := net.Listen("tcp", lis.GRPCAddress)
	defer grpcListener.Close()
	if err != nil {
		log.Fatal(err)
	}
	sftpListener, err := net.Listen("tcp", lis.SFTPAddress)
	defer sftpListener.Close()
	if err != nil {
		log.Fatal(err)
	}
	s := setup.RegisterGRPCServices()
	go func() {
		if err := s.Serve(grpcListener); err != nil {
			log.Fatal(err)
		}
		defer s.GracefulStop()
	}()

	//go func() {
	//	if err := sftpserver.ListenSFTPServer(sftpListener); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	if err := sftp.ListenSFTPServer(sftpListener); err != nil {
		log.Fatal(err)
	}
}
