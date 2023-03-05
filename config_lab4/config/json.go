package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListenPort string `json:ListenPort`
	TimeFormat string `json:TimeFormat`
}

func NewJsonConfig(configFilePath string) (Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
