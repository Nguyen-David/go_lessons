package main

import (
	"library_api/config"
	"library_api/router"
	"log"
	"os"
)


func main() {
	_, err := config.NewJsonConfig("config.json")
	if err != nil {
		log.Println(err)
		return
	}

	fileConnection, err := os.OpenFile(config.ConfigGlobal.RepoFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	defer fileConnection.Close()

	router.Route(fileConnection)
}
