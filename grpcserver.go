package main

import (
	"context"
	"io"
	"net"

	"github.com/navinds25/styx/pkg/find"
	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (s *server) FileSearch(ctx context.Context, in *pb.SearchFileInfo) (*pb.FoundFile, error) {
	return &pb.FoundFile{Type: in.Type, Match: in.Filename, Regmatch: in.Filename, Error: "nil"}, nil
}

func (s *server) FileSearchStream(stream pb.FileHandling_FileSearchStreamServer) error {
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Got %v ", in)
		dM, _, err := find.File(in.Filename)
		if err != nil {
			log.Error(err)
		}
		//dM := []string{"file1", "file2", "file3"}
		for _, name := range dM {
			f := &pb.FoundFile{}
			f.Match = name
			stream.Send(f)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %s : %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterFileHandlingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
