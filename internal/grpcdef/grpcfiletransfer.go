package grpcdef

import (
	"context"

	"github.com/navinds25/styx/pkg/sftpclient"
	log "github.com/sirupsen/logrus"

	pb "github.com/navinds25/styx/pkg/styxevent"
)

// PullFile pulls a file from a remote server
func (s *Server) PullFile(ctx context.Context, inFileCfg *pb.GetRemoteFile) (*pb.Ack, error) {
	input := &sftpclient.Input{
		Address:    inFileCfg.Sourceserver,
		Protocol:   "tcp",
		Username:   inFileCfg.Username,
		Password:   inFileCfg.Password,
		AuthMethod: inFileCfg.Authmethod,
	}
	client, err := sftpclient.CreateClient(input)
	if err != nil {
		return &pb.Ack{Message: "Error"}, err
	}
	log.Info("Created sftp connection, starting file transfer.")
	_, err = client.Pull(inFileCfg.Sourcefile, inFileCfg.Destinationfile)
	if err != nil {
		return &pb.Ack{Message: "Error"}, err
	}
	log.Info("Completed file transfer")
	return &pb.Ack{Message: "copied file from:" + inFileCfg.Sourcefile + " to:" + inFileCfg.Destinationfile}, nil
}

// ReqFileTransfer requests a file transfer on a remote node.
func (s *Server) ReqFileTransfer(ctx context.Context, inFileCfg *pb.RequestFileTransfer) (*pb.Ack, error) {
	return &pb.Ack{Message: "not implemented."}, nil
}
