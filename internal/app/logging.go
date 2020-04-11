package app

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogging() {
	logfile, err := os.OpenFile(ApplicationName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logwriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(logwriter)
	log.SetReportCaller(true)
	customLogFormat := new(log.JSONFormatter)
	customLogFormat.PrettyPrint = true
	customLogFormat.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customLogFormat)
}
