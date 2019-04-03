package app

import "github.com/takama/daemon"

// DaemonSetup sets up the daemon
func DaemonSetup() (string, error) {
	service, err := daemon.New("sftpmgmtd", "SFTP Management Daemon")
	if err != nil {
		return "", err
	}
	status, err := service.Install()
	if err != nil {
		return status, err
	}
	return status, err
}
