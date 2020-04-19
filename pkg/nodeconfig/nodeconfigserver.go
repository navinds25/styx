package nodeconfig

import (
	"context"

	pb "github.com/navinds25/styx/api/nodeconfig"
)

// Server implements the methods for the GRPC NodeConfigService
type Server struct{}

// AddNodeConfig takes a config entry and adds it to the database on all the nodes
func (ncs *Server) AddNodeConfig(ctx context.Context, in *pb.NodeConfig) (*pb.AddNodeConfigResponse, error) {
	return &pb.AddNodeConfigResponse{}, nil
}

// GetNodeConfigByID fetches NodeConfig based on node id passed.
func (ncs *Server) GetNodeConfigByID(ctx context.Context, in *pb.NodeID) (*pb.NodeConfig, error) {
	return &pb.NodeConfig{}, nil
}

// ListNodeConfig fetches all the NodeConfigs
func (ncs *Server) ListNodeConfig(ctx context.Context, in *pb.ListNodeConfigRequest) (*pb.AllNodeConfig, error) {
	return &pb.AllNodeConfig{}, nil
}

// AddNode sends a request from a new node with the nodeconfig of itself
// it returns a list of all the nodeconfigs present on the server node
func (ncs *Server) AddNode(ctx context.Context, req *pb.NodeConfig) (*pb.AllNodeConfig, error) {
	return &pb.AllNodeConfig{}, nil
}
