package domain

import "github.com/gin-gonic/gin"

// APIResponse interface for all responses
type APIResponse interface {
	ToJSON() gin.H
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (s SuccessResponse) ToJSON() gin.H {

	if s.Code == 0 {
		s.Code = 200
	}

	if s.Status == "" {
		s.Status = "success"
	}

	response := gin.H{"status": s.Status, "code": s.Code}

	if s.Data != nil {
		response["data"] = s.Data
	} else {
		if s.Message == "" {
			s.Message = "Success"
		}
		response["message"] = s.Message
	}

	return response
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (e ErrorResponse) ToJSON() gin.H {
	if e.Code == 0 {
		e.Code = 400
	}

	if e.Status == "" {
		e.Status = "error"
	}

	if e.Message == "" {
		e.Message = "An error occurred."
	}

	return gin.H{"status": e.Status, "code": e.Code, "message": e.Message}
}
