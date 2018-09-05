package main

import (
	"os"

	"github.com/navinds25/styx/pkg/sftpserver"
)

func main() {
	//defer profile.Start(profile.MemProfile).Stop()
	sftpserver.Run()
	os.Exit(1)
}
