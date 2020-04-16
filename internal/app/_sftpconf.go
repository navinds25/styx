package app

import (
	"github.com/urfave/cli"
)

var sftpConfFlagsVal sftpConfFlagsDef

type sftpConfFlagsDef struct {
	AddConfig    string
	DeleteConfig string
	ListConfig   bool
}

var sftpConfFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "add-config",
		Usage:       "provide the config to be added",
		Destination: &sftpConfFlagsVal.AddConfig,
	},
	cli.StringFlag{
		Name:        "delete-config",
		Usage:       "delete a config from the database",
		Destination: &sftpConfFlagsVal.DeleteConfig,
	},
	cli.BoolFlag{
		Name:        "list-config",
		Usage:       "lists the config from the database",
		Destination: &sftpConfFlagsVal.ListConfig,
	},
}

var sftpConfCli = cli.Command{
	Name:    "sftpconf",
	Aliases: []string{"sconf"},
	Usage:   "for the external sftp configuration",
	Action: func(c *cli.Context) error {
		Action = "sftpconf"
		return nil
	},
	Flags: append(sftpConfFlags, mainCliFlags...),
}
