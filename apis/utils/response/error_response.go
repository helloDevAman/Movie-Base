package utils_response

import "github.com/gin-gonic/gin"

type Error struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func ErrorResponse(c *gin.Context, statusCode int, r Error) {
	c.Header("Content-Type", "application/json")

	if r.Status == "" {
		r.Status = "error"
	}

	// For error response, code should not be 0
	if r.Code == 0 {
		r.Code = 1
	}

	if r.Message == "" {
		r.Message = "An error encountered at our end. Please try again later."
	}

	response := Error{
		Status:  r.Status,
		Code:    r.Code,
		Message: r.Message,
	}

	c.JSON(statusCode, response)
}
