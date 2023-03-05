package config

import (
	"encoding/json"
	"os"
)

var ConfigGlobal Config

type Config struct {
	ListeningURL string `json:ListeningURL`
	TimeFormat string `json:TimeFormat`
	Port string `json:Port`
}

func NewJsonConfig(configFilePath string) (Config, error){
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	ConfigGlobal = config

	return config, nil
}