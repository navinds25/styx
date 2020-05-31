package execute

import (
	"bufio"
	"os/exec"
	"sync"

	log "github.com/sirupsen/logrus"
)

type CmdInstance struct {
	Instance int
	Cmd      string
	Args     []string
}

type CmdMultiInstanceOutput struct {
	Instance int
	Output   string
}

type CmdError struct {
	Instance int
	Error    error
}

// BasicExecute runs a command and returns the output on completion
func BasicExecute(command string, arguments []string) (string, error) {
	program := exec.Command(command, arguments...)
	output, err := program.CombinedOutput()
	if err != nil {
		return "", err
	}
	log.Println(string(output))
	out := string(output)
	return out, nil
}

func executeAndStream(command string, arguments []string) error {
	program := exec.Command(command, arguments...)
	stdoutStream, err := program.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdoutStream)
	if err := program.Start(); err != nil {
		return err
	}
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	if err := program.Wait(); err != nil {
		return err
	}
	return nil
}

func multiExecuteAndStream(command CmdInstance) (chan CmdMultiInstanceOutput, error) {
	outputStreamChan := make(chan CmdMultiInstanceOutput)
	program := exec.Command(command.Cmd, command.Args...)
	stdoutStream, err := program.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdoutStream)
	if err := program.Start(); err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	for scanner.Scan() {
		wg.Add(1)
		go func() {
			outputStreamChan <- CmdMultiInstanceOutput{
				Instance: command.Instance, Output: scanner.Text(),
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if err := program.Wait(); err != nil {
		return nil, err
	}
	return outputStreamChan, nil
}

// MultipleExecuteAndStream executes commands currently
// TODO: function does not exit
func MultipleExecuteAndStream(commands []CmdInstance) error {
	log.Println("starting ..")
	errChan := make(chan CmdError)
	InstanceWG := sync.WaitGroup{}
	InstanceWG.Add(len(commands))
	for _, command := range commands {
		go func(command CmdInstance) {
			_, err := multiExecuteAndStream(command)
			if err != nil {
				errChan <- CmdError{
					Instance: command.Instance,
					Error:    err,
				}
			}
			InstanceWG.Done()
		}(command)
	}
	InstanceWG.Wait()
	close(errChan)
	log.Error("completed")
	return nil
}
