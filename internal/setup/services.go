package setup

import (
	ftpb "github.com/navinds25/styx/api/filetransfer"
	ncpb "github.com/navinds25/styx/api/nodeconfig"
	"github.com/navinds25/styx/pkg/filetransfer"
	"github.com/navinds25/styx/pkg/nodeconfig"
	"google.golang.org/grpc"
)

// RegisterGRPCServices returns a grpcServer with all the registered services
func RegisterGRPCServices() *grpc.Server {
	s := grpc.NewServer()
	ncpb.RegisterNodeConfigServiceServer(s, &nodeconfig.Server{})
	ftpb.RegisterFTServiceServer(s, &filetransfer.FTServer{})
	ftpb.RegisterRemoteFTServiceServer(s, &filetransfer.RemoteFTServer{})
	return s
}
