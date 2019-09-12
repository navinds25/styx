package main

import (
	"os"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/navinds25/styx/pkg/sftpclient"
	"github.com/navinds25/styx/pkg/sftpconfig"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide config as argument")
	}
	fullConf, err := sftpconfig.GetConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	conf := fullConf["sftp_config"][0]
	input := &sftpclient.Input{
		Address:    conf.RemoteHost + ":" + strconv.Itoa(conf.RemotePort),
		Protocol:   "tcp",
		Username:   conf.RemoteUser,
		Password:   conf.RemotePassword,
		AuthMethod: "pass",
		CBC:        true,
	}
	client, err := sftpclient.CreateClient(input)
	if err != nil {
		log.Fatal(err)
	}
	arrayOfFiles := client.GetListOfFiles(conf.RemotePath)
	log.Info(arrayOfFiles)

	getFilesFromArray(client, arrayOfFiles, conf.LocalPath)
	os.Exit(0)

	chanOfFiles := make(chan sftpclient.SFTPFileStat)
	for _, data := range arrayOfFiles {
		go func(data sftpclient.SFTPFileStat) {
			chanOfFiles <- data
		}(data)
	}
	//go getChannelFromArray(chanOfFiles)
	log.Println(conf.LocalPath)
	statusChan := client.MultiFilesPull(chanOfFiles, conf.LocalPath, len(arrayOfFiles))
	go func() {
		for elem := range statusChan {
			log.Println("status:", elem)
		}
		close(statusChan)
	}()
}

func getFilesFromArray(client *sftpclient.Client, arrayOfFiles []sftpclient.SFTPFileStat, outputdir string) {
	for _, file := range arrayOfFiles {
		outputfile := filepath.Join(outputdir, file.Stat.Name())
		stat, err := os.Stat(outputfile)
		if err == nil {
			clientStat, err := client.Stat(file.Path)
			if err != nil {
				log.Error(err)
				continue
			}
			if stat.Size() == clientStat.Size() {
				log.Info("File already exists:", clientStat.Name())
				continue
			} else {
				_, err := client.Pull(file.Path, outputfile)
				if err != nil {
					log.Error(err)
					continue
				}
			}
		} else if os.IsNotExist(err) {
			_, err := client.Pull(file.Path, outputfile)
			if err != nil {
				log.Error(err)
				continue
			}
		} else {
			log.Error("Error while comparing file sizes", err)
			continue
		}
	}
}
