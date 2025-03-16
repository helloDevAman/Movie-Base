package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/config"
	auth_controlers "github.com/helloDevAman/movie-base/internal/controller/auth_controllers"
	"github.com/helloDevAman/movie-base/internal/repository"
	"github.com/helloDevAman/movie-base/internal/usecase/auth_usecases"
	"github.com/helloDevAman/movie-base/utils"
)

func LoadRoutes(config *config.Config, db *sql.DB) {
	router := gin.Default()

	api := router.Group(config.APIGroup)

	// Load the auth routes
	LoadAuthRoutes(config, api, db)

	// Run the server
	router.Run(":" + config.ServerPort)
}

func LoadAuthRoutes(onfig *config.Config, api *gin.RouterGroup, db *sql.DB) {
	smsService := utils.NewTwilioService(onfig)

	authRepo := repository.NewOTPRepository()

	authRepo.InitOTPTable(db)

	authUseCase := auth_usecases.NewOTPUseCase(db, authRepo, smsService)

	authHandler := auth_controlers.OTPController{
		DB:      db,
		UseCase: authUseCase,
	}
	api.POST("/auth/send-otp", authHandler.SendOTP)
}
