package main

import (
	"log"

	"github.com/helloDevAman/movie-base/apis/config"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	log.Println("Server is running on port: ", cfg.ServerPort)
	log.Println("Connecting to DB at: ", cfg.DBHost, ":", cfg.DBPort)

}
