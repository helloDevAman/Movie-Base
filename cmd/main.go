package main

import (
	"log"

	"github.com/helloDevAman/movie-base/config"
	"github.com/helloDevAman/movie-base/internal/database"
	"github.com/helloDevAman/movie-base/internal/routes"
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

	// Initialize PostgresDatabaseConnector
	db, err := database.NewDatabaseConnector(cfg.DBType, cfg)

	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	log.Println("Connected to DB at: ", cfg.DBHost, ":", cfg.DBPort)
	defer db.Close()

	// Load the routes
	routes.LoadRoutes(cfg, db.GetConnection())

}
