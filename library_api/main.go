package main

import (
	"library_api/config"
	"library_api/router"
	"log"
	"os"
	"fmt"

	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
)


func main() {
	dsn := "host=postgres user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to database!")

	_, err = config.NewJsonConfig("config.json")
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
