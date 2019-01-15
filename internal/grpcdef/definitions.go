package grpcdef

import (
	"context"
	"io"

	"github.com/navinds25/styx/pkg/find"
	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
)

// Server struct implements
type Server struct{}

// FileSearch is for basic rpc call for searching files
func (s *Server) FileSearch(ctx context.Context, in *pb.SearchFileInfo) (*pb.FoundFile, error) {
	return &pb.FoundFile{Type: in.Type, Match: in.Filename, Regmatch: in.Filename, Error: "nil"}, nil
}

// FileSearchStream is for searching for files over an input/output grpc stream.
func (s *Server) FileSearchStream(stream pb.Styx_FileSearchStreamServer) error {
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Got ", in)
		dM, _, err := find.File(in.Filename)
		if err != nil {
			log.Error(err)
		}
		for _, name := range dM {
			f := &pb.FoundFile{}
			f.Match = name
			stream.Send(f)
		}
	}
}
