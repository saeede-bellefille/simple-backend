package main

import (
	"log"

	"github.com/saeede-bellefille/simple-backend/internal/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := api.NewServer()
	dsn := "host=localhost user=lucifer password=0990510 dbname=gotms port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	server.Setup(db)
	log.Fatal(server.Run(":8080"))
}
