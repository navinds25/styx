package styxcli

import (
	"github.com/urfave/cli"
)

var sftpConfDBFlags = []cli.Flag{}

var sftpConfDBCli = cli.Command{
	Name:    "sftpconfdb",
	Aliases: []string{"sconf"},
	Usage:   "for manipulating the configuration database for external sftp",
	Action: func(c *cli.Context) error {
		Action = "sftpconfdb"
		return nil
	},
	Flags: append(sftpConfDBFlags, mainCliFlags...),
}
