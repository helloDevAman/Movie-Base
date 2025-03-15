package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/apis/config"
	auth_routes "github.com/helloDevAman/movie-base/apis/routes/auth"
)

func LoadRoutes(config *config.Config, db *sql.DB) {
	router := gin.Default()

	api := router.Group(config.APIGroup)

	// Load the auth routes
	auth_routes.LoadAuthRoutes(api, db)

	// Run the server
	router.Run(":" + config.ServerPort)
}
