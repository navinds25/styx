package setup

import (
	executepb "github.com/navinds25/styx/api/execute"
	extensionpb "github.com/navinds25/styx/api/extension"
	ncpb "github.com/navinds25/styx/api/nodeconfig"
	"github.com/navinds25/styx/pkg/execute"
	"github.com/navinds25/styx/pkg/extension"
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
	executepb.RegisterExecuteServiceServer(s, &execute.Server{})
	extensionpb.RegisterExtensionServiceServer(s, &extension.Server{})
	//ftpb.RegisterFTServiceServer(s, &filetransfer.FTServer{})
	//ftpb.RegisterRemoteFTServiceServer(s, &filetransfer.RemoteFTServer{})
	return s, nil
}
