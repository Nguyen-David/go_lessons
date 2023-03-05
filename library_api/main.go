package main

import (
	"library_api/config"
	"library_api/router"
	"log"
)

func main() {
	_, err := config.NewJsonConfig("config.json")
	if err != nil {
		log.Println(err)
		return
	}

	router.Route()
}
