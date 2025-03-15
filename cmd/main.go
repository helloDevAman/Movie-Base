package main

import (
	"log"

	"github.com/helloDevAman/movie-base/apis/config"
)

func main() {
	cfg, error := config.LoadConfig()

	if error != nil {
		log.Fatal("Error loading .env file", error)
	}

	log.Println("Server is running on port: ", cfg.ServerPort)
	log.Println("Connecting to DB at: ", cfg.DBHost, ":", cfg.DBPort)
}
