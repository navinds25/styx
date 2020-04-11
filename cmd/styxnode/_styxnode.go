package main

import (
	"net"
	"os"
	"strconv"

	//"github.com/jasonlvhit/gocron"

	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/internal/styxcli"
	"github.com/navinds25/styx/pkg/filetransfer"
	ftpb "github.com/navinds25/styx/pkg/styxpb"
	"github.com/navinds25/styx/pkg/styxsftp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Version for inserting version via ldflags
var Version string

func runGRPCServer(lis net.Listener, s *grpc.Server) error {
	ftpb.RegisterFTServer(s, &filetransfer.FTServer{})
	log.Info("Started GRPC Server")
	if err := s.Serve(lis); err != nil {
		return err
	}
	defer s.GracefulStop()
	defer lis.Close()
	return nil
	//if err := s.Serve(lis); err != nil {
	//	return err
	//}
	//defer s.GracefulStop()
	//defer lis.Close()
	//return nil
	//pb.RegisterStyxServer(s, &grpcdef.Server{})
	//log.Info("Started GRPC Server")
	//if err := s.Serve(lis); err != nil {
	//	return err
	//}
	//defer s.GracefulStop()
	//defer lis.Close()
	//return nil
}

func main() {
	// Parse Command Line Parameters
	app.SetupLogging()
	appCli := app.Cli()
	appCli.Version = Version
	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	if err := app.MainFlagVal.CliSetDefaults(); err != nil {
		log.Fatal(err)
	}
	log.Debug("Parsed Command Line Parameters")

	// Process Application Configuration
	//config, err := styxconfig.GetConfig("_extra/config.yml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// log.Info("Processed application config")

	// Setup the Styx Databases
	//if err := app.StyxNodeDBSetup(); err != nil {
	//	log.Fatal(err)
	//}

	// Setup the SFTP Databases
	//if err := app.SFTPDBSetup(); err != nil {
	//	log.Fatal(err)
	//}
	//defer sftpdata.Data.Config.CloseConfigDB()
	//defer sftpdata.Data.Files.CloseFilesDB()
	//log.Info("Setup SFTP Databases")

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
	//go func() {
	//	log.Info("SFTP server listening on: ", sftpAddress)
	//	if err := sftpserver.ListenSFTPServer(sftplis, styxcli.MainFlagVal.SSHHOSTKEY); err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	// Run Tasks
	//scheduleI := gocron.NewScheduler()
	//scheduleI.Every(8).Seconds().Do(app.RunJobs)
	//<-scheduleI.Start()
	log.Info("SFTP server listening on: ", sftpAddress)
	if err := styxsftp.ListenSFTPServer(sftplis, styxcli.MainFlagVal.SSHHOSTKEY); err != nil {
		log.Fatal(err)
	}
}
