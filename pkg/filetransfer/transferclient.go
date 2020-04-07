package filetransfer

import (
	"context"
	"io"

	pb "github.com/navinds25/styx/pkg/filetransferpb"
	log "github.com/sirupsen/logrus"
)

// PullFiles pulls files from remote sever onto current server.
func PullFiles(ctx context.Context, ftclient pb.FTClient, remdir *pb.RemoteDirectory) error {
	stream, err := ftclient.ListDir(ctx, remdir)
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
