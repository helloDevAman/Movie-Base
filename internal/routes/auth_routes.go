package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/internal/controllers"
)

func LoadAuthRoutes(api *gin.RouterGroup, db *sql.DB) {
	api.GET("/auth/verify-mobile", controllers.VerifyMobileRequest(db))
}
