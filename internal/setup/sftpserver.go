package setup

import (
	"net"

	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/sftp"
)

// ServeSFTPServer is the main function for the Styx SFTP Server
func ServeSFTPServer(lis net.Listener) error {
	if err := sftp.ListenSFTPServer(lis, app.MainFlagVal.SSHHOSTKEY); err != nil {
		return err
	}
	return nil
}
