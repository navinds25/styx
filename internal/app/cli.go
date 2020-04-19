package app

import (
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// ApplicationName is the name of the app for the CLI.
const ApplicationName = "styx"

// MainFlagVal is an instance of struct for the main cli flags
var MainFlagVal MainFlags

// Action is for the cli Commands
var Action string

// SubAction is for the subcommands in the cli
var SubAction string

// MainFlags is a struct for the main cli flags
type MainFlags struct {
	SetupService        bool
	DataDir             string
	Debug               bool
	Help                bool
	Version             bool
	InterfaceAddress    string
	GrpcPort            int
	SftpPort            int
	SSHHOSTKEY          string
	Config              string
	OverwriteHostConfig bool
	EncryptionKey       string
	MasterIP            string
	MasterPort          int
}

var mainCliFlags = []cli.Flag{
	cli.BoolFlag{
		Name:        "setup-service",
		Destination: &MainFlagVal.SetupService,
	},
	cli.StringFlag{
		Name:        "datadir",
		Usage:       "provide the data dir",
		Destination: &MainFlagVal.DataDir,
	},
	cli.BoolFlag{
		Name:        "debug",
		Destination: &MainFlagVal.Debug,
	},
	cli.StringFlag{
		Name:        "address",
		Value:       "0.0.0.0",
		Destination: &MainFlagVal.InterfaceAddress,
	},
	cli.IntFlag{
		Name:        "grpcport",
		Value:       28888,
		Destination: &MainFlagVal.GrpcPort,
	},
	cli.IntFlag{
		Name:        "sftpport",
		Value:       28889,
		Destination: &MainFlagVal.SftpPort,
	},
	cli.StringFlag{
		Name:        "i",
		Value:       "ssh_host_rsa_key",
		Destination: &MainFlagVal.SSHHOSTKEY,
	},
	cli.StringFlag{
		Name:        "config",
		Value:       "nodeconfig.yml",
		Destination: &MainFlagVal.Config,
	},
	cli.BoolFlag{
		Name:        "overwrite-hostconfig",
		Destination: &MainFlagVal.OverwriteHostConfig,
	},
	cli.StringFlag{
		Name:        "encryption-key",
		Destination: &MainFlagVal.EncryptionKey,
	},
	cli.StringFlag{
		Name:        "master-ip",
		Value:       "127.0.0.1",
		Destination: &MainFlagVal.MasterIP,
	},
	cli.IntFlag{
		Name:        "master-port",
		Value:       28888,
		Destination: &MainFlagVal.MasterPort,
	},
}

// CliSetDefaults sets default values for empty cli flags.
func (cliflags *MainFlags) CliSetDefaults() error {
	if MainFlagVal.Help || MainFlagVal.Version {
		os.Exit(0)
	}
	if MainFlagVal.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logs enabled!")
	}

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	if cliflags.DataDir == "" {
		cliflags.DataDir = path.Join(currentDir, "data")
	}
	return nil
}

// Cli for all commandline arguments.
func Cli() *cli.App {
	app := cli.NewApp()
	cli.HelpFlag = cli.BoolFlag{
		Name:        "h",
		Destination: &MainFlagVal.Help,
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:        "v",
		Destination: &MainFlagVal.Version,
	}
	app.Name = ApplicationName
	app.Usage = "The file transfer application"
	app.Commands = []cli.Command{
		//sftpConfCli,
	}
	app.Flags = mainCliFlags
	return app
}
