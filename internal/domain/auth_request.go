package domain

type VerifyMobileRequest struct {
	Mobile string `json:"mobile" binding:"required"`
}
