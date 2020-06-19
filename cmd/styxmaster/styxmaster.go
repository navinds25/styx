package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	executepb "github.com/navinds25/styx/api/execute"
)

// Version for inserting version via ldflags
var Version string

func main() {
	//setup.Logging()
	//appCli := app.Cli()
	//appCli.Version = Version
	//if err := appCli.Run(os.Args); err != nil {
	//	log.Fatal(err)
	//}
	//s, err := setup.MasterSetup()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//grpcListener, err := net.Listen("tcp", lis.GRPCAddress)
	//defer grpcListener.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	address := "localhost:28888"
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := executepb.NewExecuteServiceClient(conn)
	argInDesign := &executepb.Executable{
		Command:   os.Args[1],
		Arguments: []string{"-port", "10001"},
	}
	log.Println(argInDesign)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	output, err := c.Run(ctx, argInDesign)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(output)
}
