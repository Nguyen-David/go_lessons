package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YmlConfig struct {
	ListenPort string `yaml:"ListenPort"`
	TimeFormat string `yaml:"TimeFormat"`
}

func NewYmlConfig(configFilePath string) (YmlConfig, error) {
	yfile, err := os.ReadFile(configFilePath)
	if err != nil {
		return YmlConfig{}, err
	}

	var config YmlConfig
	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return YmlConfig{}, err
	}

	return config, nil
}
