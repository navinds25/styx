package main

import (
	"net"
	"os"

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
	hcM, lis, err := setup.NodeSetup()
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
	setup.Services(grpcListener, sftpListener, hcM)
}
