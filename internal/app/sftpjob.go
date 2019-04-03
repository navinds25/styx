package app

import (
	"github.com/navinds25/styx/pkg/sftpdata"
	log "github.com/sirupsen/logrus"
)

// RunJobs running the configured jobs via gocron.
func RunJobs() error {
	allConfigs, err := sftpdata.Data.Config.GetAll()
	if err != nil {
		return err
	}
	log.Debug("All configs", allConfigs)
	for i, config := range allConfigs {
		log.Println("processing config: ", i, config)
	}
	return nil
}
