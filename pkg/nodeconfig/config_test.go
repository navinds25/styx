package nodeconfig_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/navinds25/styx/pkg/nodeconfig"
)

func TestConfigFromYAML(t *testing.T) {
	type args struct {
		inputYaml io.Reader
	}
	type test struct {
		name    string
		args    args
		want    *nodeconfig.ConfigInput
		wantErr bool
	}
	tests := []test{}
	positive_args := args{
		inputYaml: bytes.NewReader([]byte(`
host_config:
  node_id: node1
  node_type: internal
  ip_address: 127.0.0.1
  grpc_port: 28888
  sftp_port: 28889
  sz: app
  sftp_auth:
    hostkey_file: ssh_host_rsa_key #need to fix this field
  external_access: true

master_config:
  master_ip: 127.0.0.1
  master_port: 28888
`)),
	}
	positive_want := &nodeconfig.ConfigInput{
		HostConfig: &nodeconfig.HostConfigInput{
			NodeID:    "node1",
			NodeType:  "internal",
			IPAddress: "127.0.0.1",
			GRPCPort:  28888,
			SFTPPort:  28889,
			SZ:        "app",
			SFTPAuth: nodeconfig.SFTPAuth{
				HostkeyFile: "ssh_host_rsa_key",
			},
			ExternalAccess: true,
		},
		MasterConfig: &nodeconfig.MasterConfigInput{
			MasterIP: "127.0.0.1", MasterPort: 28888,
		},
	}
	positive := test{
		name: "positive", args: positive_args, want: positive_want, wantErr: false,
	}
	tests = append(tests, positive)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nodeconfig.ConfigFromYAML(tt.args.inputYaml)
			t.Log("got args: ", tt.args.inputYaml)
			t.Logf("processed:\nhostconfig: %+v , \nmaster config: %+v", got.HostConfig, got.MasterConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ConfigFromYAML() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
