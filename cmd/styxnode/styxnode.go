package main

import (
	"io"
	"net"
	"os"
	"strconv"

	"github.com/jasonlvhit/gocron"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/internal/grpcdef"
	"github.com/navinds25/styx/internal/styxcli"
	"github.com/navinds25/styx/pkg/sftpdata"
	"github.com/navinds25/styx/pkg/sftpserver"
	pb "github.com/navinds25/styx/pkg/styxevent"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Version for inserting version via ldflags
var Version string

func runGRPCServer(lis net.Listener, s *grpc.Server) error {
	pb.RegisterStyxServer(s, &grpcdef.Server{})
	log.Info("Started GRPC Server")
	if err := s.Serve(lis); err != nil {
		return err
	}
	defer s.GracefulStop()
	defer lis.Close()
	return nil
}

func init() {
	logfile, err := os.OpenFile(styxcli.ApplicationName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logwriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(logwriter)
	log.SetReportCaller(true)
	customLogFormat := new(logrus.JSONFormatter)
	customLogFormat.PrettyPrint = true
	customLogFormat.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customLogFormat)
}

func main() {
	// Parse Command Line Parameters
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

	// Process Application Configuration
	//config, err := styxconfig.GetConfig("_extra/config.yml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// log.Info("Processed application config")

	// Setup Databases
	if err := app.DBSetup(); err != nil {
		log.Fatal(err)
	}
	defer sftpdata.Data.Config.CloseConfigDB()
	defer sftpdata.Data.Files.CloseFilesDB()
	log.Info("Setup Databases")

	// GRPC Server
	grpcAddress := styxcli.MainFlagVal.InterfaceAddress + ":" + strconv.Itoa(styxcli.MainFlagVal.GrpcPort)
	grpclis, err := net.Listen("tcp", grpcAddress)
	defer grpclis.Close()
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	go func() {
		log.Info("GRPC server listening on: ", grpcAddress)
		if err := runGRPCServer(grpclis, s); err != nil {
			log.Fatal(err)
		}
	}()

	// SFTP Server
	sftpAddress := styxcli.MainFlagVal.InterfaceAddress + ":" + strconv.Itoa(styxcli.MainFlagVal.SftpPort)
	sftplis, err := net.Listen("tcp", sftpAddress)
	defer sftplis.Close()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Info("SFTP server listening on: ", sftpAddress)
		if err := sftpserver.ListenSFTPServer(sftplis, styxcli.MainFlagVal.SSHHOSTKEY); err != nil {
			log.Fatal(err)
		}
	}()

	// Run Tasks
	scheduleI := gocron.NewScheduler()
	scheduleI.Every(8).Seconds().Do(app.RunJobs)
	<-scheduleI.Start()
}
