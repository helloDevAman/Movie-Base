package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/config"
)

func LoadRoutes(config *config.Config, db *sql.DB) {
	router := gin.Default()

	api := router.Group(config.APIGroup)

	// Load the auth routes
	LoadAuthRoutes(api, db)

	// Run the server
	router.Run(":" + config.ServerPort)
}
