package app

import (
	"log"

	"github.com/helloDevAman/movie-base/config"
	"github.com/helloDevAman/movie-base/pkg/database"
	routes "github.com/helloDevAman/movie-base/pkg/route"
)

func Run(cfg *config.Config) {
	// Database
	log.Println("Connecting to DB at: ", cfg.DB.Host, ":", cfg.DB.Port)
	pgConnector := database.ConnectNewPostgresDB(cfg)

	db, err := pgConnector.Connect()

	if err != nil {
		log.Printf("Error: %v", err)
	}
	defer db.Close()

	// Start the server
	ginRouteLoader := routes.LoadNewGinRoute(cfg, db)
	router, err := ginRouteLoader.LoadRoutes()
	if err != nil {
		log.Fatalf("Failed to load routes: %v", err)
	}

	if err := ginRouteLoader.StartListening(router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
