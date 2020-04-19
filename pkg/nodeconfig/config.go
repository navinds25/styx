package nodeconfig

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// Input is an interface for  Config/Cli Input
type Input interface {
	GetInputFromConfig(string)
}

// ConfigInput is the main struct for all configuration
type ConfigInput struct {
	HostConfig   *HostConfigInput   `json:"host_config,omitempty"`
	MasterConfig *MasterConfigInput `json:"master_config,omitempty"`
}

// ConfigFromYAML returns a Config from yaml
func ConfigFromYAML(inputYaml io.Reader) (*ConfigInput, error) {
	config := &ConfigInput{}
	data, err := ioutil.ReadAll(inputYaml)
	if err != nil {
		return nil, err
	}
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonData, config); err != nil {
		return nil, err
	}
	return config, nil
}

// ConfigFromYAMLFile returns a ConfigInput from a Yaml File
func ConfigFromYAMLFile(filename string) (*ConfigInput, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	fileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return ConfigFromYAML(fileReader)
}

// GetInputFromConfig retrieves the config from a config file
func (ci *ConfigInput) GetInputFromConfig(filename string) error {
	ci, err := ConfigFromYAMLFile(filename)
	if err != nil {
		return err
	}
	return nil
}
