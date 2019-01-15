package main

import (
	"github.com/navinds25/styx/pkg/sftpclient"
	log "github.com/sirupsen/logrus"
)

func main() {
	i := sftpclient.Input{}
	i.Address = "127.0.0.1:28888"
	i.AuthMethod = "pass"
	i.Protocol = "tcp"
	i.Password = "tiger"
	i.Username = "testusr"
	client, err := sftpclient.CreateClient(i)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := client.Conn.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	dirInfo, err := client.Conn.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range dirInfo {
		log.Println(entry.Name())
	}
	client.Conn.Close()
}
