package auth_usecases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/helloDevAman/movie-base/internal/domain"
	"github.com/helloDevAman/movie-base/internal/repository"
	"github.com/helloDevAman/movie-base/utils"
)

type OTPUseCase struct {
	DB         *sql.DB
	OTPRepo    repository.OTPRepository
	SMSService utils.SMSService
}

func NewOTPUseCase(db *sql.DB, repo repository.OTPRepository, smsService utils.SMSService) *OTPUseCase {
	return &OTPUseCase{DB: db, OTPRepo: repo, SMSService: smsService}
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

	err = u.SMSService.SendOTP(mobile, otpCode)
	if err != nil {
		log.Println("Error sending OTP: ", err)
		return nil, fmt.Errorf("failed to send OTP via Twilio")
	}

	return otp, nil
}
