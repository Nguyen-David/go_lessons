package main

import (
	"config_lab4/config"
	"fmt"
	"log"
)

func main() {
	conf, err := config.NewJsonConfig("config.json")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(conf.ListenPort)
	fmt.Println(conf.TimeFormat)

	yconf, err := config.NewYmlConfig("config.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(yconf.ListenPort)
	fmt.Println(yconf.TimeFormat)

	envconf, err := config.NewEnvConfig("config.env")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(envconf.ListenPort)
	fmt.Println(envconf.TimeFormat)

	flagconf, err := config.NewFlagConfig()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(flagconf.ListenPort)
	fmt.Println(flagconf.TimeFormat)

}
