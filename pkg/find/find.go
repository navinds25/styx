package find

import (
	"regexp"

	"github.com/kr/fs"
)

// File is used for finding a file.
func File(filename string) ([]string, []string, error) {
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
