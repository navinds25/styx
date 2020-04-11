package nodeconfig

import (
	"context"

	pb "github.com/navinds25/styx/pkg/styxpb"
)

// NodeConfigServer implements the methods for the GRPC NodeConfigService
type NodeConfigServer struct{}

// AddNodeConfig takes a config entry and adds it to the database on all the nodes
func (ncs *NodeConfigServer) AddNodeConfig(ctx context.Context, in *pb.NodeConfig) (*pb.AddNodeConfigResponse, error) {
	return &pb.AddNodeConfigResponse{}, nil
}

// GetNodeConfigByID fetches NodeConfig based on node id passed.
func (ncs *NodeConfigServer) GetNodeConfigByID(ctx context.Context, in *pb.NodeID) (*pb.NodeConfig, error) {
	return &pb.NodeConfig{}, nil
}

// ListNodeConfig fetches all the NodeConfigs
func (ncs *NodeConfigServer) ListNodeConfig(ctx context.Context, in *pb.ListNodeConfigRequest) (*pb.AllNodeConfig, error) {
	return &pb.AllNodeConfig{}, nil
}
