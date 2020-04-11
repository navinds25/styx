package main

import (
	"os"

	"github.com/navinds25/styx/internal/app"
	log "github.com/sirupsen/logrus"
)

// Version for inserting version via ldflags
var Version string

func main() {
	app.SetupLogging()
	appCli := app.Cli()
	appCli.Version = Version
	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	if err := app.SetupNode(); err != nil {
		log.Fatal(err)
	}
}
