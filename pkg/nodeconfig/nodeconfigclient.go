package nodeconfig

import (
	"context"
	"fmt"

	pb "github.com/navinds25/styx/api/nodeconfig"
	"google.golang.org/grpc"
)

// HostConfigToNodeConfig takes a HostConfig and returns a NodeConfig
func HostConfigToNodeConfig(hcM *HostConfigModel) *pb.NodeConfig {
	nodetype := pb.NodeConfig_NODETYPE_NODE
	sz := pb.NodeConfig_SZ_UNSPECIFIED
	out := &pb.NodeConfig{
		NodeId:      hcM.NodeID,
		Nodetype:    nodetype,
		GrpcAddress: hcM.GRPCAddress,
		SftpAddress: hcM.SFTPAddress,
		Sz:          sz,
		GrpcAuth: &pb.GRPCAuth{
			TlsCert: hcM.GRPCAuth.TLSCertFile,
		},
	}
	return out
}

// AddNodeClient send the calling nodes config to the server
// and gets all the config from the server
func AddNodeClient(conn *grpc.ClientConn) error {
	hcM, err := Data.NodeConfig.GetHostConfigEntry("hostconfig")
	nc := HostConfigToNodeConfig(hcM)
	client := pb.NewNodeConfigServiceClient(conn)
	allnc, err := client.AddNode(context.TODO(), nc)
	if err != nil {
		return err
	}
	fmt.Println(allnc)
	return nil
}
