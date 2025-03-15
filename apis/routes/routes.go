package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/apis/config"
)

func LoadRoutes(config *config.Config) {
	router := gin.Default()

	api := router.Group(config.APIGroup)

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Routes loaded successfully: ", router.Routes())

	router.Run(":" + config.ServerPort)
}
