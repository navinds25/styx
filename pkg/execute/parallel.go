package execute

import (
	"bufio"
	"os/exec"
	"sync"

	log "github.com/sirupsen/logrus"
)

// CommandRequest is for type request in the queue
type CommandRequest struct {
	ID        string
	Command   string
	Arguments []string
	Instance  int
	Port      int
}

// CommandQueue is the queue for the Command instances
var CommandQueue = make(chan CommandRequest)

// CommandQuitRequest is for stopping a Command Instance
type CommandQuitRequest struct {
	ID string
}

// CommandQuitChannel closes all the commands
var CommandQuitChannel = make(chan CommandQuitRequest)

// SubscribeCmdAsyncParallel subscribes to the CommandQueue and starts a command parallelly when present
func SubscribeCmdAsyncParallel() {
	go func() {
		for {
			cmdInst := <-CommandQueue
			program := exec.Command(cmdInst.Command, cmdInst.Arguments...)
			stdoutStream, err := program.StdoutPipe()
			if err != nil {
				log.Error("Error setting up reading stdout of program: ", err)
			}
			scanner := bufio.NewScanner(stdoutStream)
			if err := program.Start(); err != nil {
				log.Error("Error starting program", err)
			}
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				for scanner.Scan() {
					log.WithField("instance", cmdInst.Instance).WithField("port", cmdInst.Port).Println(scanner.Text())
				}
				defer wg.Done()
			}()
			select {
			case cmdInst := <-CommandQueue:
				program := exec.Command(cmdInst.Command, cmdInst.Arguments...)
				stdoutStream, err := program.StdoutPipe()
				if err != nil {
					log.Error("Error setting up reading stdout of program: ", err)
				}
				scanner := bufio.NewScanner(stdoutStream)
				if err := program.Start(); err != nil {
					log.Error("Error starting program", err)
				}
				wg := sync.WaitGroup{}
				wg.Add(1)
				go func() {
					for scanner.Scan() {
						log.WithField("instance", cmdInst.Instance).WithField("port", cmdInst.Port).Info(scanner.Text())
					}
					defer wg.Done()
				}()
				wg.Wait()
				if err := program.Wait(); err != nil {
					log.Error("Error waiting for program to complete: ", err)
				}
			case cmdQuitChan := <-CommandQuitChannel:
				if cmdQuitChan.ID == cmdInst.ID {
					if err := program.Process.Kill(); err != nil {
						log.Error("Error killing program: ", err)
					}
					log.Infof("Killed Command ID %s on request", cmdInst.ID)
				}
			}
		}
	}()
}
