package execute

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func basicExecute(command string, arguments []string) (string, error) {
	program := exec.Command(command, arguments...)
	output, err := program.CombinedOutput()
	if err != nil {
		return "", err
	}
	log.Println(string(output))
	out := string(output)
	return out, nil
}
