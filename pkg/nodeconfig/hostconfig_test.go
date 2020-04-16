package nodeconfig_test

import (
	"io"
	"reflect"
	"testing"

	"github.com/navinds25/styx/pkg/nodeconfig"
)

// Test Function YamlToHostConfig

func TestHostConfigFromYAML(t *testing.T) {
	type args struct {
		inputYaml io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *nodeconfig.HostConfigInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nodeconfig.HostConfigFromYAML(tt.args.inputYaml)
			if (err != nil) != tt.wantErr {
				t.Errorf("HostConfigFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HostConfigFromYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
