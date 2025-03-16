package utils

import (
	"log"

	"github.com/helloDevAman/movie-base/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

// SMSService is the interface for sending SMS
type SMSService interface {
	SendOTP(to string, otp string) error
}

// TwilioService is the implementation of SMSService for Twilio
type TwilioService struct {
	client          *twilio.RestClient
	twilioServiceID string
}

// NewTwilioService creates a new TwilioService
func NewTwilioService(cfg *config.Config) *TwilioService {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TWILIOAccountSID,
		Password: cfg.TWILIOAuthToken,
	})

	return &TwilioService{
		client:          client,
		twilioServiceID: cfg.TWILIOApiSecret,
	}
}

// SendOTP sends an OTP using Twilio
func (s *TwilioService) SendOTP(to string, otp string) error {

	log.Println("Sending OTP: ", otp, "to: ", to)
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	_, err := s.client.VerifyV2.CreateVerification(s.twilioServiceID, params)

	if err != nil {
		return err
	}

	return nil
}
