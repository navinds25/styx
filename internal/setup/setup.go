package setup

import (
	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/nodeconfig"
	log "github.com/sirupsen/logrus"
)

// NodeSetup is the main setup function
func NodeSetup() (*nodeconfig.HostConfigModel, error) {
	// read cli flags
	if err := app.MainFlagVal.CliSetDefaults(); err != nil {
		return nil, err
	}
	// setup dbs
	if err := DBsetup(app.MainFlagVal.DataDir); err != nil {
		return nil, err
	}
	log.Debug("completed the db setup")
	// get the hostconfig
	hcM, err := updateHostConfig()
	if err != nil {
		log.Error("error updating hostconfig: ", err)
		return nil, err
	}
	QueueSetup()
	return hcM, nil
}
