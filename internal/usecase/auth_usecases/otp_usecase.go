package auth_usecases

import (
	"database/sql"
	"errors"
	"time"

	"github.com/helloDevAman/movie-base/internal/domain"
	"github.com/helloDevAman/movie-base/internal/repository"
	"github.com/helloDevAman/movie-base/utils"
)

type OTPUseCase struct {
	DB      *sql.DB
	OTPRepo repository.OTPRepository
}

func NewOTPUseCase(db *sql.DB, repo repository.OTPRepository) *OTPUseCase {
	return &OTPUseCase{DB: db, OTPRepo: repo}
}

func (u *OTPUseCase) SendOTP(db *sql.DB, mobile string) (*domain.OTP, error) {
	otpCode := utils.GenerateOTP(6)
	expiry := time.Now().Add(5 * time.Minute)

	otp := &domain.OTP{
		Mobile:    mobile,
		Code:      otpCode,
		ExpiresAt: expiry,
	}

	err := u.OTPRepo.SaveOTP(db, otp)
	if err != nil {
		return nil, err
	}

	// Simulate sending OTP (In production, integrate Twilio or other services)
	// pkg.SendSMSTwilio(mobile, otpCode)

	return otp, nil
}

func (u *OTPUseCase) VerifyOTP(db *sql.DB, mobile, code string) error {
	otp, err := u.OTPRepo.GetOTP(db, mobile)
	if err != nil {
		return errors.New("OTP not found")
	}

	if time.Now().After(otp.ExpiresAt) {
		return errors.New("OTP expired")
	}

	if otp.Code != code {
		return errors.New("Invalid OTP")
	}

	return nil
}
