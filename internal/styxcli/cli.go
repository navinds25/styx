package styxcli

import (
	"os"
	"path"
	"path/filepath"

	"github.com/urfave/cli"
)

// ApplicationName for the name of the application.
const ApplicationName = "styx"

// MainFlagVal is an instance of struct for the main cli flags
var MainFlagVal MainFlags

// Action is for the cli Commands
var Action string

// SubAction is for the subcommands in the cli
var SubAction string

// MainFlags is a struct for the main cli flags
type MainFlags struct {
	SetupService     bool
	DataDir          string
	Debug            bool
	Help             bool
	Version          bool
	InterfaceAddress string
	GrpcPort         int
	SftpPort         int
	SSHHOSTKEY       string
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
		Value:       28889,
		Destination: &MainFlagVal.GrpcPort,
	},
	cli.IntFlag{
		Name:        "sftpport",
		Value:       28888,
		Destination: &MainFlagVal.SftpPort,
	},
	cli.StringFlag{
		Name:        "i",
		Value:       "ssh_host_rsa_key",
		Destination: &MainFlagVal.SSHHOSTKEY,
	},
}

// GetCliFlags is a Factory function returning the struct for MainCli
func (cliflags *MainFlags) GetCliFlags() error {
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
		sftpConfCli,
	}
	app.Flags = mainCliFlags
	return app
}
