package helper

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a 6-digit OTP
func GenerateOTP() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return fmt.Sprintf("%06d", r.Intn(1000000))
}
