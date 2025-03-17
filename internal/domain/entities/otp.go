package entities

import "time"

type OTP struct {
	MobileNumber string
	Code         string
	CreatedAt    time.Time
	ExpiresAt    time.Time
}
