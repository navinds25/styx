package sftpclient

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// BasicCopyfromRemote copies a file from a remote sftp server.
func BasicCopyfromRemote(inputfile, outputfile string, client *Client) error {
	input, err := client.Conn.Open(inputfile)
	if err != nil {
		return err
	}
	output, err := os.Create(outputfile)
	if err != nil {
		log.Error("Error creating output file: %s ; Err: %v", outputfile, err)
		return err
	}
	bytes, err := io.Copy(output, input)
	if err != nil {
		log.Error("Error copying file %s to %s, Err: %v", inputfile, outputfile, err)
		return err
	}
	defer func() {
		if err := input.Close(); err != nil {
			log.Error("Error closing input file: ", inputfile, err)
		}
	}()
	fmt.Printf("%d bytes copied\n", bytes)
	if err = output.Sync(); err != nil {
		log.Error("Error syncing changes to disk", err)
		return err
	}
	defer func() {
		if err := output.Close(); err != nil {
			log.Error(err)
		}
	}()
	return nil
}

// Pull copies inputfile from a remote sftp server as outputfile on the local filesystem.
func Pull(inputfile, outputfile string, client *Client) (int64, error) {
	input, err := client.Conn.Open(inputfile)
	if err != nil {
		log.Error("Error creating input file: ", inputfile)
		return 0, err
	}
	defer input.Close()
	output, err := os.Create(outputfile)
	if err != nil {
		log.Error("Error creating output file: ", outputfile)
		return 0, err
	}
	defer output.Close()
	bytesTransferred, err := io.Copy(output, input)
	if err != nil {
		log.Error("Error while transferring file")
		return bytesTransferred, err
	}
	return bytesTransferred, nil
}
