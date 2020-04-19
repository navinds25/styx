package app

//var nodeConfigCli = cli.Command{
//	Name:    "nodeconfig",
//	Aliases: []string{"hostconfig", "nodeconf"},
//	Usage:   "for the node configuration",
//	Action: func(c *cli.Context) error {
//		Action = "nodeconfig"
//		return nil
//	},
//	Flags: append(nodeConfigCliFlags, mainCliFlags...),
//}
//
//var nodeConfigCliFlags = []cli.Flag{
//	cli.BoolFlag{
//		Name:        "setup-service",
//		Destination: &MainFlagVal.SetupService,
//	},
//	cli.StringFlag{
//		Name:        "datadir",
//		Usage:       "provide the data dir",
//		Destination: &MainFlagVal.DataDir,
//	},
//	cli.BoolFlag{
//		Name:        "debug",
//		Destination: &MainFlagVal.Debug,
//	},
//	cli.StringFlag{
//		Name:        "address",
//		Value:       "0.0.0.0",
//		Destination: &MainFlagVal.InterfaceAddress,
//	},
//	cli.IntFlag{
//		Name:        "grpcport",
//		Value:       28888,
//		Destination: &MainFlagVal.GrpcPort,
//	},
//	cli.IntFlag{
//		Name:        "sftpport",
//		Value:       28889,
//		Destination: &MainFlagVal.SftpPort,
//	},
//	cli.StringFlag{
//		Name:        "i",
//		Value:       "ssh_host_rsa_key",
//		Destination: &MainFlagVal.SSHHOSTKEY,
//	},
//	cli.StringFlag{
//		Name:        "config",
//		Value:       "nodeconfig.yml",
//		Destination: &MainFlagVal.Config,
//	},
//	cli.BoolFlag{
//		Name:        "overwrite-hostconfig",
//		Destination: &MainFlagVal.OverwriteHostConfig,
//	},
//	cli.StringFlag{
//		Name:        "encryption-key",
//		Destination: &MainFlagVal.EncryptionKey,
//	},
//	cli.StringFlag{
//		Name:        "master-ip",
//		Value:       "127.0.0.1",
//		Destination: &MainFlagVal.MasterIP,
//	},
//	cli.IntFlag{
//		Name:        "master-port",
//		Value:       28888,
//		Destination: &MainFlagVal.MasterPort,
//	},
//}
//
//// NodeConfigFlagVal is an instance of struct for the nodeconfig cli flags
//var NodeConfigFlagVal nodeconfig.ConfigInput
//
