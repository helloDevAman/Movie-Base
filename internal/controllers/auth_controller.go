package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/internal/domain"
	"github.com/helloDevAman/movie-base/internal/utils"
)

func VerifyMobileRequest(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		data := map[string]string{
			"mobile": "1234567890",
		}

		response := domain.SuccessResponse{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Mobile number verified successfully",
			Data:    data,
		}

		utils.SendResponse(c, http.StatusOK, response)
	}
}
