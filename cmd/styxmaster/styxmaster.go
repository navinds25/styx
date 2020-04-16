package main

import (
	"log"
	"os"

	"github.com/navinds25/styx/internal/app"

	"github.com/navinds25/styx/internal/setup"
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
	//s, err := setup.MasterSetup()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//grpcListener, err := net.Listen("tcp", lis.GRPCAddress)
	//defer grpcListener.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
