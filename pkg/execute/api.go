package execute

import (
	"context"

	pb "github.com/navinds25/styx/api/execute"
)

// Server implements all the grpc server methods for the extension package
type Server struct{}

// Run is grpc server method for calling the specified command with arguments
func (ps *Server) Run(ctx context.Context, in *pb.Executable) (*pb.Output, error) {
	output, err := basicExecute(in.Command, in.Arguments)
	errString := func(err error) string {
		if err == nil {
			return ""
		}
		return err.Error()
	}(err)
	return &pb.Output{Output: output, Error: errString}, nil
}
