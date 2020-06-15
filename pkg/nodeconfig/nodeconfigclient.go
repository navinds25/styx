package nodeconfig

import (
	"context"
	"crypto/x509"

	pb "github.com/navinds25/styx/api/nodeconfig"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	log.Infof("Got list of nodeconfigs: %+v", allnc)
	return nil
}

// GRPCDialerMaster returns a grpc connection to the master
func GRPCDialerMaster() (*grpc.ClientConn, error) {
	mc, err := Data.NodeConfig.GetMasterConfigEntry()
	if err != nil {
		return nil, err
	}
	hcM, err := Data.NodeConfig.GetHostConfigEntry(HostConfigKey)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	status := certPool.AppendCertsFromPEM(hcM.GRPCAuth.TLSCertBinData)
	if !status {
		return nil, err
	}
	conn, err := grpc.Dial(mc.Address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, "")))
	return conn, nil
}

// GRPCDialerNode returns a grpc connection to the nodeid and type specified
func GRPCDialerNode(nodeID, nodeType string) error {
	return nil
}

// UpdateNodeConfigClient gets the list of nodeconfigs from the master and updates the db
func UpdateNodeConfigClient() (*pb.AllNodeConfig, error) {
	conn, err := GRPCDialerMaster()
	if err != nil {
		return nil, err
	}
	client := pb.NewNodeConfigServiceClient(conn)
	allConfig, err := client.ListNodeConfig(context.TODO(), &pb.NodeConfigRequest{Prefix: NodeConfigPrefixKey})
	if err != nil {
		return nil, err
	}
	for _, config := range allConfig.AllNodeConfig {
		key := NodeConfigPrefixKey + DBSeparator + config.Nodetype.String() + DBSeparator + config.NodeId
		value := config
		if err := Data.NodeConfig.AddNodeConfigEntry(key, value); err != nil {
			return allConfig, err
		}
	}
	return allConfig, nil
}
