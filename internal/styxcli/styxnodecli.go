package styxcli

import "github.com/urfave/cli"

// InterfaceAddress for cli input for listen address
var InterfaceAddress = "0.0.0.0"

// GrpcPort for cli input for grpc port
var GrpcPort = "28889"

// SftpPort for cli input for sftp port
var SftpPort = "28888"

// SSHHOSTKEY for cli input for ssh host key
var SSHHOSTKEY = "ssh_host_rsa_key"

// App returns an cli app object for running the cli
func App() *cli.App {
	app := cli.NewApp()
	app.Name = "styx"
	app.Usage = "For styx"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "address",
			Value:       "0.0.0.0",
			Destination: &InterfaceAddress,
		},
		cli.StringFlag{
			Name:        "grpcport",
			Value:       "28889",
			Destination: &GrpcPort,
		},
		cli.StringFlag{
			Name:        "sftpport",
			Value:       "28888",
			Destination: &SftpPort,
		},
		cli.StringFlag{
			Name:        "i",
			Value:       "ssh_host_rsa_key",
			Destination: &SSHHOSTKEY,
		},
	}
	return app
}
