package domain

import "time"

// OTP entity
type OTP struct {
	Mobile    string    `json:"mobile"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}
