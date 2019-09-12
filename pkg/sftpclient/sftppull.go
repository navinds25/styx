package sftpclient

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/kr/fs"
	log "github.com/sirupsen/logrus"
)

// SFTPFileStat is for file info on a remote sftp file.
type SFTPFileStat struct {
	Path string
	Stat os.FileInfo
	Err  error
}

// Pull copies inputfile from a remote sftp server as outputfile on the local filesystem.
func (client *Client) Pull(inputfile, outputfile string) (int64, error) {
	input, err := client.Open(inputfile)
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

// Open is a wrapper around sftp Open function
func (client *Client) Open(filepath string) (io.ReadWriteCloser, error) {
	if client.CBC {
		return client.CBCConn.Open(filepath)
	}
	return client.Conn.Open(filepath)
}

func (client *Client) multiFilePull(wg *sync.WaitGroup, file SFTPFileStat, outputFile string) error {
	defer wg.Done()
	input, err := client.Open(file.Path)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		return err
	}
	return nil
}

// MultiFilesPull is for pulling multiple files in parallel
func (client *Client) MultiFilesPull(fileList chan SFTPFileStat, outputPath string, length int) chan SFTPFileStat {
	var wg sync.WaitGroup
	errChan := make(chan SFTPFileStat)
	wg.Add(length)
	defer wg.Wait()
	go func() {
		for file := range fileList {
			outputFile := filepath.Join(outputPath, file.Stat.Name())
			if err := client.multiFilePull(&wg, file, outputFile); err != nil {
				file.Err = err
				errChan <- file
			} else {
				errChan <- file
			}
		}
	}()
	return errChan
}

// Walk is a wrapper around the sftp walk function
func (client *Client) Walk(rootDir string) *fs.Walker {
	if client.CBC {
		return client.CBCConn.Walk(rootDir)
	}
	return client.Conn.Walk(rootDir)
}

// GetListOfFiles gets the list of files on a remote sftp server
func (client *Client) GetListOfFiles(rootDir string) []SFTPFileStat {
	arrayOfFiles := []SFTPFileStat{}
	walker := client.Walk(rootDir)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}
		stat := walker.Stat()
		if stat.IsDir() {
			continue
		} else {
			arrayOfFiles = append(arrayOfFiles, SFTPFileStat{
				Stat: stat,
				Path: walker.Path(),
			})
		}
	}
	return arrayOfFiles
}

// Stat is a wrapper around the the sftp Stat function
func (client *Client) Stat(filename string) (os.FileInfo, error) {
	if client.CBC {
		return client.CBCConn.Stat(filename)
	}
	return client.Conn.Stat(filename)
}
