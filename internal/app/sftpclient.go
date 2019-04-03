package app

import (
	"github.com/navinds25/styx/pkg/sftpclient"
	log "github.com/sirupsen/logrus"
)

// RunSftpClient runs sftpclient for something
func RunSftpClient(inFile, outFile string, pull bool) error {
	config := sftpclient.Input{
		Address:    "127.0.0.1:28888",
		Protocol:   "tcp",
		Username:   "testusr",
		Password:   "tiger",
		AuthMethod: "pass",
	}

	client, err := sftpclient.CreateClient(&config)
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
	bytesTransferred, err := client.Push(inFile, "write/"+outFile)
	if err != nil {
		log.Error(err)
	}
	log.Println(bytesTransferred)

	defer client.Conn.Close()

	return nil
}
