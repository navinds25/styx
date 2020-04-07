package main

import (
	"context"
	"io"
	"os"
	"time"

	//"github.com/jasonlvhit/gocron"

	"github.com/navinds25/styx/internal/styxcli"
	"github.com/navinds25/styx/pkg/filetransfer"
	ftpb "github.com/navinds25/styx/pkg/filetransferpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var Version string

func init() {
	logfile, err := os.OpenFile(styxcli.ApplicationName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logwriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(logwriter)
	log.SetReportCaller(true)
	customLogFormat := new(log.JSONFormatter)
	customLogFormat.PrettyPrint = true
	customLogFormat.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customLogFormat)
}

func main() {
	// process cli
	appCli := styxcli.Cli()
	appCli.Version = Version
	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	if err := styxcli.MainFlagVal.GetCliFlags(); err != nil {
		log.Fatal(err)
	}
	if styxcli.MainFlagVal.Help || styxcli.MainFlagVal.Version {
		os.Exit(0)
	}
	if styxcli.MainFlagVal.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logs enabled!")
	}
	log.Info("Parsed Command Line Parameters")

	// open config database
	// process configuration file
	// send identity to styxnodes
	// send config to target styxnodes

	conn, err := grpc.Dial("127.0.0.1:28889", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	ftclient := ftpb.NewFTClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filetransfer.ListFiles(ctx, ftclient, &ftpb.RemoteDirectory{SourcePath: "testdata"})
	filetransfer.ListFilesCondition(ctx, ftclient)
}
