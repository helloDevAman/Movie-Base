package usecase

import (
	"fmt"
	"time"

	"github.com/helloDevAman/movie-base/internal/domain/entities"
	"github.com/helloDevAman/movie-base/internal/repository"
	"github.com/helloDevAman/movie-base/pkg/helper"
)

type SendOTPUseCase struct {
	repo repository.OTPRepository
}

func NewSendOTPUseCase(repo repository.OTPRepository) *SendOTPUseCase {
	return &SendOTPUseCase{repo: repo}
}

func (u *SendOTPUseCase) Execute(mobileNumber string) (*entities.OTP, error) {
	latestOTP, err := u.repo.GetLatestOTP(mobileNumber)
	if err == nil && time.Since(latestOTP.CreatedAt) < 60*time.Second {
		remaining := 60 - int(time.Since(latestOTP.CreatedAt).Seconds())
		return nil, fmt.Errorf("try after %d seconds", remaining)
	}

	otpCode := helper.GenerateOTP()
	otp := entities.OTP{
		MobileNumber: mobileNumber,
		Code:         otpCode,
		CreatedAt:    time.Now(),
	}

	if err := u.repo.SaveOTP(otp); err != nil {
		return nil, fmt.Errorf("failed to save OTP: %w", err)
	}

	// In real-world, send OTP via SMS service.
	fmt.Printf("Generated OTP for %s: %s\n", mobileNumber, otpCode)
	return &otp, nil
}
