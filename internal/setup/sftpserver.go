package setup

import (
	"net"

	"github.com/navinds25/styx/pkg/nodeconfig"
	"github.com/navinds25/styx/pkg/sftp"
	log "github.com/sirupsen/logrus"
)

// ServeSFTPServer is the main function for the Styx SFTP Server
func ServeSFTPServer(lis net.Listener) error {
	hcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
	if err != nil {
		log.Error("error getting hostkkey from db")
		return err
	}
	if err := sftp.ListenSFTPServer(lis, hcM.SFTPAuth.HostkeyFile); err != nil {
		return err
	}
	return nil
}
