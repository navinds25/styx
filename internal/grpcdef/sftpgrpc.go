package grpcdef

import (
	"context"

	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
)

// GRPCTest is the function for testing grpc connectivity.
func (s *Server) GRPCTest(ctx context.Context, in *pb.Ack) (*pb.Ack, error) {
	log.Info(in.Message)
	return &pb.Ack{
		Message: "received message.",
	}, nil
}
