package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/internal/domain/validation"
	"github.com/helloDevAman/movie-base/internal/usecase"
	message "github.com/helloDevAman/movie-base/pkg/massage"
)

type SendOTPHandler struct {
	useCase *usecase.SendOTPUseCase
}

func NewSendOTPHandler(useCase *usecase.SendOTPUseCase) *SendOTPHandler {
	return &SendOTPHandler{useCase: useCase}
}

type SendOTPRequest struct {
	MobileNumber string `json:"mobile" binding:"required"`
}

func (h *SendOTPHandler) ServeHTTP(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": message.MethodNotAllowed})
		return
	}

	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": message.InvalidRequestBody})
		return
	}

	if err := validation.ValidateMobileNumber(req.MobileNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	otp, err := h.useCase.Execute(req.MobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message.OTPSentSuccessfully, "otp": otp.Code})
}
