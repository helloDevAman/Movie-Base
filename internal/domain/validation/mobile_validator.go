package validation

import (
	"errors"
	"log"
	"regexp"
)

func ValidateMobileNumber(mobileNumber string) error {
	re := regexp.MustCompile(`^\+91\d{10}$`)
	if !re.MatchString(mobileNumber) {
		log.Println("Mobile:", mobileNumber)
		return errors.New("invalid mobile number format")
	}
	return nil
}
