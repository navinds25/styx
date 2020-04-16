package setup

import (
	ncpb "github.com/navinds25/styx/api/nodeconfig"
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/nodeconfig"
	"google.golang.org/grpc"
)

// MasterSetup is the main setup func for the styxmaster
func MasterSetup() (*grpc.Server, error) {
	// read cli flags
	if err := app.MainFlagVal.CliSetDefaults(); err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	ncpb.RegisterNodeConfigServiceServer(s, &nodeconfig.Server{})
	return s, nil
}
