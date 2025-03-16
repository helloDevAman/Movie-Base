package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/config"
	auth_controlers "github.com/helloDevAman/movie-base/internal/controller/auth_controllers"
	"github.com/helloDevAman/movie-base/internal/repository"
	"github.com/helloDevAman/movie-base/internal/usecase/auth_usecases"
)

func LoadRoutes(config *config.Config, db *sql.DB) {
	router := gin.Default()

	api := router.Group(config.APIGroup)

	// Load the auth routes
	LoadAuthRoutes(api, db)

	// Run the server
	router.Run(":" + config.ServerPort)
}

func LoadAuthRoutes(api *gin.RouterGroup, db *sql.DB) {
	authRepo := repository.NewOTPRepo()

	authRepo.InitOTPTable(db)

	authUseCase := auth_usecases.NewOTPUseCase(db, authRepo)

	authHandler := auth_controlers.OTPHandler{
		DB:      db,
		UseCase: authUseCase,
	}
	api.POST("/auth/send-otp", authHandler.SendOTP)
}
