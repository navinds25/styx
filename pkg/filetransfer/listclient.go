package filetransfer

import (
	"context"
	"io"
	"time"

	log "github.com/sirupsen/logrus"

	pb "github.com/navinds25/styx/api/filetransfer"
)

// ListFiles lists files on Remote Server
func ListFiles(ctx context.Context, ftclient pb.FTClient, remdir *pb.RemoteDirectory) error {
	stream, err := ftclient.ListDir(ctx, remdir)
	if err != nil {
		return err
	}
	for {
		file, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Info(file)
	}
	return nil
}

// ListFilesCondition - needs input args refactoring required.
func ListFilesCondition(ctx context.Context, ftclient pb.FTClient) error {
	duration, err := time.ParseDuration("1.5h")
	if err != nil {
		return err
	}
	startTime := time.Now().Unix() - int64(duration.Seconds())
	timeval := &pb.TimeValues{
		TimeStart: startTime,
	}
	log.Println(startTime)
	conditiontr := &pb.Condition{
		ConditionType: pb.Condition_Time,
		TimeValues:    timeval,
	}
	conditions := []*pb.Condition{}
	conditions = append(conditions, conditiontr)
	remdir := &pb.RemoteDirectoryCondition{
		SourcePath: "testdata/time",
		Condition:  conditions,
	}
	stream, err := ftclient.ListDirCondition(ctx, remdir)
	if err != nil {
		return err
	}
	for {
		feature, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		log.Println(feature)
	}
	return nil
}
