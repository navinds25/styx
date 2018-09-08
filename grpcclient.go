package main

import (
	"context"
	"io"
	"os"
	"time"

	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const address = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect ", err)
	}
	defer conn.Close()
	c := pb.NewFileHandlingClient(conn)
	clientDeadline := time.Now().Add(time.Duration(10000) * time.Millisecond)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	p := &pb.SearchFileInfo{}
	p.Filename = "hello.txt"
	p.Type = "filesearch"
	p.Server = "local"
	stream, err := c.FileSearchStream(ctx)
	if err != nil {
		log.Fatal("Problem executing grpc func")
	}
	if err := stream.Send(p); err != nil {
		log.Error("Error sending message from client")
		os.Exit(1)
	}
	if err := stream.CloseSend(); err != nil {
		log.Error("Error closing grpc client sending stream")
		os.Exit(1)
	}
	for {
		file, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Error(err)
				os.Exit(1)
			}
		}
		log.Println(file)
	}
}
