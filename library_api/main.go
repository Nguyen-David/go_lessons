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
	host := os.Getenv("POSTGRES_DB_HOST")
	user := os.Getenv("POSTGRES_DB_USER")
	password := os.Getenv("POSTGRES_DB_PASSWORD")
	db := os.Getenv("POSTGRES_DB_NAME")
	port := os.Getenv("POSTGRES_DB_PORT")
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, db, port)
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
