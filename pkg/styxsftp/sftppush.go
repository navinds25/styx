package styxsftp

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Push copies inputfile to a remote sftp server as outputfile
func (client *Client) Push(inputfile, outputfile string) (int64, error) {
	output, err := client.Conn.Create(outputfile)
	if err != nil {
		log.Error("Error with Output file")
		return 0, err
	}
	defer output.Close()
	input, err := os.Open(inputfile)
	if err != nil {
		log.Error("Error with Input file")
		return 0, err
	}
	defer input.Close()

	bytesTransfered, err := io.Copy(output, input)
	if err != nil {
		log.Error("Error while copying")
		return bytesTransfered, err
	}
	return bytesTransfered, nil
}
