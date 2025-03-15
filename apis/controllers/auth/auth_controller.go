package auth_controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	utils_response "github.com/helloDevAman/movie-base/apis/utils/response"
)

func VerifyMobileRequest(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// data := map[string]string{"is_send": "true"}
		utils_response.ErrorResponse(c, http.StatusOK, utils_response.Error{Message: "Verify Mobile Request", Code: 0})
	}
}
