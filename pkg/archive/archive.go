package archive

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
)

// CreateArchive has the params for creating archives
type CreateArchive struct {
	InputFile   string
	OutputFile  string
	Dirs        bool
	Directories []string
	Type        string
}

// GetBaseName returns the base name of a file, i.e. without extension
// eg: GetBaseName(blah.test.csv) returns: blah.test
func GetBaseName(inputFilename string) string {
	out1 := strings.Split(inputFilename, ".")
	out2 := out1[:len(out1)-1]
	outputFilename := strings.Join(out2, ".")
	return outputFilename
}

func createArchive(input []string, outputfile, archiveType string) error {
	switch archiveType {
	case "tar.bz2":
		if err := archiver.TarBz2.Make(outputfile, input); err != nil {
			return err
		}
	case "tar.gz":
		if err := archiver.TarGz.Make(outputfile, input); err != nil {
			return err
		}
	case "zip":
		if err := archiver.Zip.Make(outputfile, input); err != nil {
			return err
		}
	default:
		return fmt.Errorf("archive type: %s, is not valid please use: tar.bz2, tar.gz or zip", archiveType)
	}
	return nil
}

// Create creates an archive using params in archive.CreateArchive
func Create(params CreateArchive) error {
	if params.Type == "" {
		log.Info("Using default compression tar.bz2")
		params.Type = "tar.bz2"
	}
	if params.Dirs {
		if params.Directories == nil {
			return errors.New("Dirs bool specified, but no Directories specified")
		}
		if params.OutputFile == "" {
			return errors.New("Output file not specified")
		}
	}
	if !params.Dirs {
		if params.InputFile == "" {
			return errors.New("Input file should be specified or Dirs bool should be true")
		}
		if params.OutputFile == "" {
			params.OutputFile = GetBaseName(params.InputFile) + params.Type
		}
		params.Directories = append(params.Directories, params.InputFile)
	}

	err := createArchive(params.Directories, params.OutputFile, params.Type)
	if err != nil {
		return err
	}
	return nil
}
