package extension

import (
	"context"
	"errors"
	"strings"

	pb "github.com/navinds25/styx/api/extension"
	log "github.com/sirupsen/logrus"
)

// Server implements all the grpc server methods for the extension package
type Server struct{}

// ExecuteExtension is grpc server method for calling the specified extension
func (ps *Server) ExecuteExtension(ctx context.Context, in *pb.ExecuteExtensionParams) (*pb.ExecuteExtensionOutput, error) {
	extension := Directory[strings.ToUpper(in.Name)]
	if extension == nil {
		log.Error("got extension name: ", in.Name)
		return nil, errors.New("No valid extension name specified")
	}
	log.Debug("running function now")
	output, err := extension.Run(in.Arguments)
	errString := func(err error) string {
		if err == nil {
			return ""
		}
		return err.Error()
	}
	log.Debug(output)
	output["error"] = errString(err)
	return &pb.ExecuteExtensionOutput{
		Output: output,
	}, nil
}
