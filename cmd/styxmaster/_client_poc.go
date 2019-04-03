package main

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/navinds25/styx/pkg/sftpclient"
	pb "github.com/navinds25/styx/pkg/styxevent"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const address = "127.0.0.1:28889"

func testSFTP() {
	i := &sftpclient.Input{
		Address:    "127.0.0.1:28888",
		AuthMethod: "pass",
		Protocol:   "tcp",
		Password:   "tiger",
		Username:   "testusr",
	}
	client, err := sftpclient.CreateClient(i)
	if err != nil {
		log.Error(err)
	}
	dir, err := client.Conn.Getwd()
	if err != nil {
		log.Error(err)
	}
	log.Println(dir)
	dirInfo, err := client.Conn.ReadDir(dir)
	if err != nil {
		log.Error(err)
	}
	for _, entry := range dirInfo {
		log.Println(entry.Name())
	}
	client.Conn.Close()
}

func streammain() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect ", err)
	}
	defer conn.Close()
	c := pb.NewStyxClient(conn)
	clientDeadline := time.Now().Add(time.Duration(1000000) * time.Millisecond)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	p := &pb.SearchFileInfo{
		Filename: "hello.txt",
		Type:     "filesearch",
		Server:   "local",
	}
	stream, err := c.FileSearchStream(ctx)
	if err != nil {
		log.Fatal(err)
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

func grpcmain() {
	conn, err := grpc.Dial("0.0.0.0:8432", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := pb.NewSftpClient(conn)
	deadline := time.Now().Add(time.Duration(1000000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	transferTypeValue := pb.SftpTransferConfig_TransferType_value["Pull"]
	transferType := pb.SftpTransferConfig_TransferType(transferTypeValue)
	msg, err := c.AddConfig(ctx, &pb.SftpTransferConfig{
		Transferid:     "test1",
		Type:           transferType,
		Remoteuser:     "testuser",
		Remotepassword: "tiger",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(msg)
	msg, err = c.GRPCTest(ctx, &pb.Ack{Message: "sending this message"})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(msg)
}

func main() {
	testSFTP()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect ", err)
	}
	defer conn.Close()
	c := pb.NewStyxClient(conn)
	clientDeadline := time.Now().Add(time.Duration(1000000) * time.Millisecond)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	p := &pb.GetRemoteFile{
		Type:            "filepull",
		Sourcefile:      "/home/sysusr/install.sh",
		Destinationfile: "/home/navin/deldir/install.sh",
		Sourceserver:    "13.233.92.228:28888",
		Authmethod:      "pass",
		Username:        "testusr",
		Password:        "tiger",
	}
	log.Println("Sending message to pull file")
	message, err := c.PullFile(ctx, p)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(message)
}
