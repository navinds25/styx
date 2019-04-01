package grpcdef

import (
	"context"
	"log"

	"github.com/navinds25/styx/pkg/sftpconfig"
	"github.com/navinds25/styx/pkg/sftpdata"
	pb "github.com/navinds25/styx/pkg/styxevent"
)

// AddConfig adds a config to config db.
func (s *Server) AddConfig(ctx context.Context, tc *pb.SftpTransferConfig) (*pb.Ack, error) {
	transferConfig := &sftpconfig.TransferConfig{
		TransferID:     tc.Transferid,
		Description:    tc.Description,
		Type:           tc.Type.String(),
		LocalFile:      tc.Localfile,
		LocalPath:      tc.Localpath,
		RemoteFile:     tc.Remotefile,
		RemotePath:     tc.Remotepath,
		RemoteHost:     tc.Remotehost,
		RemotePort:     int(tc.Remoteport),
		RemoteUser:     tc.Remoteuser,
		RemotePassword: tc.Remotepassword,
	}
	byteArray := []byte(transferConfig.RemoteUser)
	log.Println("size ", len(byteArray))
	if err := sftpdata.Data.Config.AddSFTPEntry(transferConfig); err != nil {
		return &pb.Ack{
			Error: err.Error(),
		}, err
	}
	return &pb.Ack{
		Message: "Added config successfully.",
	}, nil
}

// ListConfig lists the configs from the db.
func (s *Server) ListConfig(_ *pb.Noparams, _ pb.Styx_ListConfigServer) error {
	return nil
}
