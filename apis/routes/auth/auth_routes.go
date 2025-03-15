package auth_routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	auth_controllers "github.com/helloDevAman/movie-base/apis/controllers/auth"
)

func LoadAuthRoutes(api *gin.RouterGroup, db *sql.DB) {
	api.GET("/auth/verify-mobile", auth_controllers.VerifyMobileRequest(db))
}
