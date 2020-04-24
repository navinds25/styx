package setup

import (
	ncpb "github.com/navinds25/styx/api/nodeconfig"
	"github.com/navinds25/styx/pkg/nodeconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// RegisterGRPCServices returns a grpcServer with all the registered services
func RegisterGRPCServices(grpcAuth *nodeconfig.GRPCAuthModel) (*grpc.Server, error) {
	creds, err := credentials.NewServerTLSFromFile(grpcAuth.TLSCertFile, grpcAuth.TLSKeyFile)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer(grpc.Creds(creds))
	ncpb.RegisterNodeConfigServiceServer(s, &nodeconfig.Server{})
	//ftpb.RegisterFTServiceServer(s, &filetransfer.FTServer{})
	//ftpb.RegisterRemoteFTServiceServer(s, &filetransfer.RemoteFTServer{})
	return s, nil
}
