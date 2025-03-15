package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/internal/domain"
)

func SendResponse(c *gin.Context, statusCode int, response domain.APIResponse) {
	c.JSON(statusCode, response.ToJSON())
}
