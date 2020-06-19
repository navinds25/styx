package setup

import (
	"crypto/x509"
	"errors"
	"io/ioutil"
	"os"

	"github.com/navinds25/styx/internal/app"
	"github.com/navinds25/styx/pkg/nodeconfig"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// logic for this is weird but correct as it accounts for default value in Cli.
func readHostConfigFromCli() (*nodeconfig.HostConfigInput, error) {
	if app.MainFlagVal.Config != "" {
		hcInput, err := nodeconfig.HostConfigFromYAMLFile(app.MainFlagVal.Config)
		if err != nil {
			return nil, err
		}
		return hcInput, nil
	}
	return nil, nil
}

func readConfigFromCli() (*nodeconfig.ConfigInput, error) {
	if app.MainFlagVal.Config != "" {
		input, err := nodeconfig.ConfigFromYAMLFile(app.MainFlagVal.Config)
		if err != nil {
			return nil, err
		}
		return input, nil
	}
	return nil, nil
}

func overwriteFromCli(cliHC *nodeconfig.HostConfigInput) (*nodeconfig.HostConfigModel, error) {
	if cliHC == nil {
		return nil, errors.New("config didn't have any value")
	}
	hcM, masterModel, err := nodeconfig.HostConfigToModel(cliHC)
	if err != nil {
		log.Debug("conversion of cli input to hostconfig model failed with ", err)
		return nil, err
	}
	if err := nodeconfig.Data.NodeConfig.AddHostConfigEntry(nodeconfig.HostConfigKey, hcM); err != nil {
		log.Debug("failed to add hostconfig entry to database ", err)
		return nil, err
	}
	log.Debug("added entry for hostconfig")
	if err := nodeconfig.Data.NodeConfig.AddMasterConfigEntry(masterModel); err != nil {
		log.Debug("failed to add entry for master config")
		return nil, err
	}
	dHcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
	if err != nil {
		log.Debug("failed to get hostconfig entry from database ", err)
		return dHcM, err
	}
	log.Debug("got entry from DB", dHcM)
	return dHcM, nil
}

func addNode(certFile string) error {
	mc, err := nodeconfig.Data.NodeConfig.GetMasterConfigEntry()
	if err != nil {
		return err
	}
	certPool := x509.NewCertPool()
	certFD, err := os.Open(certFile)
	if err != nil {
		log.Error("error opening tlscert", err)
		return err
	}
	cert, err := ioutil.ReadAll(certFD)
	if err != nil {
		log.Error("error reading tlscert: ", err)
		return err
	}
	status := certPool.AppendCertsFromPEM(cert)
	if !status {
		log.Error("certificate pool could not be updated")
		return err
	}
	conn, err := grpc.Dial(mc.Address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, "")))
	if err != nil {
		log.Error("could not send request to master")
		return err
	}
	defer conn.Close()
	if err := nodeconfig.AddNodeClient(conn); err != nil {
		log.Error("error from AddNodeClient", err)
		return err
	}

	return nil
}

func updateHostConfig() (*nodeconfig.HostConfigModel, error) {
	if !app.MainFlagVal.OverwriteHostConfig {
		hcM, err := nodeconfig.Data.NodeConfig.GetHostConfigEntry(nodeconfig.HostConfigKey)
		if err != nil {
			log.Error("error reading hostconfig, overwrite is not passed: ", err)
			return nil, err
		}
		log.Debugf("got hostconfig from db: %v", hcM)
		return hcM, nil
	}
	log.Debug("reading config from cli")
	config, err := readConfigFromCli()
	if err != nil {
		log.Error("error reading config", err)
		return nil, err
	}
	log.Debugf("got config: %+v", config.HostConfig)
	hcM, err := overwriteFromCli(config.HostConfig)
	if err != nil {
		log.Error("error updating hostconfig: ", err)
		return nil, err
	}
	log.Debugf("got hostconfig from db/overwrite: %v", hcM)
	if err := addNode(hcM.GRPCAuth.TLSCertFile); err != nil {
		log.Error("error adding node to cluster", err)
		//return nil, err
	}
	return hcM, nil
}
