package main

import (
	"regexp"

	"github.com/kr/fs"
	log "github.com/sirupsen/logrus"
)

func main() {
	dM, rM, err := findFile("list_of_files")
	if err != nil {
		log.Error(err)
	}
	log.Println("Directly matched files:", dM)
	log.Println("Regex Matched files:", rM)
}

func findFile(filename string) ([]string, []string, error) {
	r, err := regexp.Compile(filename)
	if err != nil {
		return nil, nil, err
	}
	directMatches := []string{}
	regexMatches := []string{}
	walker := fs.Walk("/")
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}
		stat := walker.Stat()
		if stat.IsDir() {
			continue
		} else {
			regmatch := r.FindString(stat.Name())
			if stat.Name() == filename {
				path := walker.Path()
				directMatches = append(directMatches, path)
			}
			if regmatch != "" {
				path := walker.Path()
				regexMatches = append(regexMatches, path)
			}
		}
	}
	return directMatches, regexMatches, nil
}
