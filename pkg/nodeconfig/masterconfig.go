package nodeconfig

// MasterConfigInput is the
type MasterConfigInput struct {
	MasterIP      string `json:"master_ip,omitempty"`
	GRPCPort      int    `json:"grpc_port,omitempty"`
	MasterAddress string `json:"master_address,omitempty"`
}
