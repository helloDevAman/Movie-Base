package utils_response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, r Response) {
	c.Header("Content-Type", "application/json")

	if r.Status == "" {
		r.Status = "success"
	}

	if r.Data != nil && r.Message != "" {
		r.Message = ""
	}

	response := Response{
		Status:  r.Status,
		Code:    r.Code,
		Message: r.Message,
		Data:    r.Data,
	}

	c.JSON(statusCode, response)
}
