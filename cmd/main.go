package main

import (
	"log"

	"github.com/helloDevAman/movie-base/apis/config"
)

func main() {
	// Load the configuration
	loader := &config.EnvConfigLoader{}
	cfg := config.LoadEnvConfig(loader)
	if cfg == nil {
		log.Fatalf("Failed to load configuration")
	}

	log.Println("Server is running on port: ", cfg.ServerPort)
	log.Println("Connecting to DB at: ", cfg.DBHost, ":", cfg.DBPort)

}
