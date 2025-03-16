package auth_controlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/internal/domain"
	"github.com/helloDevAman/movie-base/internal/usecase/auth_usecases"
	"github.com/helloDevAman/movie-base/utils"
)

type OTPHandler struct {
	DB      *sql.DB
	UseCase *auth_usecases.OTPUseCase
}

func (h *OTPHandler) SendOTP(c *gin.Context) {
	var request struct {
		Mobile string `json:"mobile" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, domain.ErrorResponse{
			Message: "Mobile number is required",
		})
		return
	}

	// Call the UseCase function to generate and store OTP
	otp, err := h.UseCase.SendOTP(h.DB, request.Mobile)
	if err != nil {
		log.Println("Error is", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	// Return the response
	utils.SendResponse(c, http.StatusOK, domain.SuccessResponse{
		Data: gin.H{"message": "OTP sent successfully", "otp": otp.Code},
	})
}
