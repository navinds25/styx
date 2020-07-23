package setup

import (
	indesignpb "github.com/navinds25/EviveInDesignServer/api/indesign"
	indesign "github.com/navinds25/EviveInDesignServer/pkg/indesign"
	ftpb "github.com/navinds25/styx/api/filetransfer"
	ncpb "github.com/navinds25/styx/api/nodeconfig"
	"github.com/navinds25/styx/pkg/filetransfer"
	"github.com/navinds25/styx/pkg/nodeconfig"
	"google.golang.org/grpc"
)

// RegisterGRPCServices returns a grpcServer with all the registered services
func RegisterGRPCServices(grpcAuth *nodeconfig.GRPCAuthModel) (*grpc.Server, error) {
	//creds, err := credentials.NewServerTLSFromFile(grpcAuth.TLSCertFile, grpcAuth.TLSKeyFile)
	//if err != nil {
	//	return nil, err
	//}
	//s := grpc.NewServer(grpc.Creds(creds))
	s := grpc.NewServer()
	ncpb.RegisterNodeConfigServiceServer(s, &nodeconfig.Server{})
	ftpb.RegisterFTServiceServer(s, &filetransfer.FTServer{})
	ftpb.RegisterRemoteFTServiceServer(s, &filetransfer.RemoteFTServer{})
	// external grpc services:
	indesignpb.RegisterInDesignServiceServer(s, &indesign.Server{})
	return s, nil
}
