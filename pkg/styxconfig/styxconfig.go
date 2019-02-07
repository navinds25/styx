package styxconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"

	myaml "github.com/ghodss/yaml"
)

// TransferConfig is the configuration for an individual transfer
type TransferConfig struct {
	TransferID        string `json:"transfer_id"`
	SourceFile        string `json:"source_file"`
	SourceServer      string `json:"source_server"`
	DestinationFile   string `json:"destination_file"`
	DestinationServer string `json:"destination_server"`
	Credentials       struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}

// Config is Main Styx Application Config
type Config struct {
	Servers            map[string]string `json:"servers"`
	ScheduledTransfers []TransferConfig  `json:"scheduled_transfers"`
}

// GetConfig returns the main application yaml config
func GetConfig(configFile string) (*Config, error) {
	var config Config
	_, err := os.Stat(configFile)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	jsonData, err := myaml.YAMLToJSON(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, err
	}
	//if err := yaml.NewDecoder(file).Decode(&config); err != nil {
	//	return nil, err
	//}
	return &config, nil
}
