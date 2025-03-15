package models

type VerifyMobileRequest struct {
	Mobile string `json:"mobile" binding:"required"`
}
