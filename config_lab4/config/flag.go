package config

import (
	"flag"
)

type FlagConfig struct {
	ListenPort string
	TimeFormat string
}

func NewFlagConfig() (FlagConfig, error) {
	ListenPort := flag.String("listen_port", "", "a string")
	TimeFormat := flag.String("time_format", "", "a string")
	flag.Parse()

	config := FlagConfig{
		ListenPort: *ListenPort,
		TimeFormat: *TimeFormat,
	}

	return config, nil
}
