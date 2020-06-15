package nodeconfig

import (
	"context"

	pb "github.com/navinds25/styx/api/nodeconfig"
	log "github.com/sirupsen/logrus"
)

// Server implements the methods for the GRPC NodeConfigService
type Server struct{}

// AddNodeConfig takes a config entry and adds it to the database on all the nodes
func (ncs *Server) AddNodeConfig(ctx context.Context, in *pb.NodeConfig) (*pb.AddNodeConfigResponse, error) {
	return &pb.AddNodeConfigResponse{}, nil
}

// GetNodeConfigByID fetches NodeConfig based on node id passed.
func (ncs *Server) GetNodeConfigByID(ctx context.Context, in *pb.GetNodeRequest) (*pb.NodeConfig, error) {
	return &pb.NodeConfig{}, nil
}

// ListNodeConfig fetches all the NodeConfigs
func (ncs *Server) ListNodeConfig(ctx context.Context, in *pb.NodeConfigRequest) (*pb.AllNodeConfig, error) {
	data, err := Data.NodeConfig.GetAllNodeConfigEntries(in.Prefix)
	if err != nil {
		log.Errorf("Error getting all nodeconfigs for key prefix %s, err: %s", NodeConfigPrefixKey, err.Error())
	}
	return &pb.AllNodeConfig{
		AllNodeConfig: data,
	}, nil
}

// UpdateAllNodeConfig reads the list of NodeConfigs from the master and updates the local node.
// returns all the configuration added in the database
func (ncs *Server) UpdateAllNodeConfig(ctx context.Context, in *pb.NodeConfigRequest) (*pb.AllNodeConfig, error) {
	allConfig, err := UpdateNodeConfigClient()
	if err != nil {
		log.Error(err)
	}
	return allConfig, nil
}

// AddNode sends a request from a new node with the nodeconfig of itself
// it returns a list of all the nodeconfigs present on the server node
func (ncs *Server) AddNode(ctx context.Context, req *pb.NodeConfig) (*pb.AllNodeConfig, error) {
	key := NodeConfigPrefixKey + DBSeparator + req.Nodetype.String() + DBSeparator + req.NodeId
	log.Debug("Adding nodeconfig with key", key)
	if err := Data.NodeConfig.AddNodeConfigEntry(key, req); err != nil {
		log.Errorf("Error adding nodeconfig to db with key %s and value %+v , err: %s", key, req, err.Error())
	}
	data, err := Data.NodeConfig.GetAllNodeConfigEntries(NodeConfigPrefixKey)
	if err != nil {
		log.Errorf("Error getting all nodeconfigs for key prefix %s, err: %s", NodeConfigPrefixKey, err.Error())
	}
	return &pb.AllNodeConfig{
		AllNodeConfig: data,
	}, nil
}
