package main

import (
	"net"
	"os"

	"github.com/navinds25/EviveInDesignServer/pkg/indesign"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/internal/setup"
	"github.com/navinds25/styx/pkg/nodeconfig"
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
	hcM, err := setup.NodeSetup()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("main function, hcM val: ", hcM)
	log.Debug("grpcAddress: ", hcM.GRPCAddress)
	log.Debug("sftpAddress: ", hcM.SFTPAddress)
	defer nodeconfig.Data.NodeConfig.CloseDB()
	defer indesign.Data.InDesign.Close()
	grpcListener, err := net.Listen("tcp", hcM.GRPCAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcListener.Close()
	sftpListener, err := net.Listen("tcp", hcM.SFTPAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpListener.Close()
	s, err := setup.RegisterGRPCServices(hcM.GRPCAuth)
	if err != nil {
		log.Fatal(err)
	}
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
	if err := setup.ServeSFTPServer(sftpListener); err != nil {
		log.Fatal(err)
	}
}
