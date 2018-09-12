package app

import (
	"github.com/navinds25/styx/pkg/sftpclient"
	log "github.com/sirupsen/logrus"
)

// RunSftpClient runs sftpclient for something
func RunSftpClient(inFile, outFile string, pull bool) error {
	config := sftpclient.Input{
		Address:    "35.200.177.113:26773",
		Protocol:   "tcp",
		Username:   "navin",
		Password:   "xi6XTtpk6q7nw8oaL6DLfgd",
		PrivateKey: "/home/navin/.ssh/google_compute_engine",
		AuthMethod: "pk",
	}

	client, err := sftpclient.CreateClient(config)
	if err != nil {
		return err
	}

	cwd, err := client.Conn.Getwd()
	if err != nil {
		return err
	}

	files, err := client.Conn.ReadDir(cwd)
	if err != nil {
		return err
	}
	for _, file := range files {
		log.Println(file.Name())
	}
	bytesTransferred, err := sftpclient.Push(inFile, "write/"+outFile, client)
	if err != nil {
		log.Error(err)
	}
	log.Println(bytesTransferred)

	defer client.Conn.Close()

	return nil
}
