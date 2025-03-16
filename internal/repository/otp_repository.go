package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/helloDevAman/movie-base/internal/domain"
)

type OTPRepository interface {
	InitOTPTable(db *sql.DB) error
	SaveOTP(db *sql.DB, otp *domain.OTP) error
	GetOTP(db *sql.DB, mobile string) (*domain.OTP, error)
}

type OTPRepositoryImpl struct{}

func NewOTPRepository() OTPRepository {
	return &OTPRepositoryImpl{}
}

// Initialize OTP table if not exists
func (r *OTPRepositoryImpl) InitOTPTable(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS otps (
    mobile VARCHAR(15) PRIMARY KEY,
    code VARCHAR(6) NOT NULL,
    expires_at TIMESTAMP NOT NULL
	)`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Failed to create OTP table", err)
		return fmt.Errorf("failed to create OTP table: %v", err)
	}
	fmt.Println("OTP table is ready!")
	return nil
}

// Insert or update OTP in the database
func (r *OTPRepositoryImpl) SaveOTP(db *sql.DB, otp *domain.OTP) error {

	upsertQuery := `
	INSERT INTO otps (mobile, code, expires_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (mobile) DO UPDATE SET
		code = EXCLUDED.code,
		expires_at = EXCLUDED.expires_at;
	`

	_, err := db.Exec(upsertQuery, otp.Mobile, otp.Code, otp.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to insert/update OTP: %v", err)
	}

	return nil
}

// Get OTP by mobile number
func (r *OTPRepositoryImpl) GetOTP(db *sql.DB, mobile string) (*domain.OTP, error) {
	// Query OTP from the database
	query := `
	SELECT mobile, code, expires_at
	FROM otps
	WHERE mobile = ?
	`

	row := db.QueryRow(query, mobile)

	otp := &domain.OTP{}
	err := row.Scan(&otp.Mobile, &otp.Code, &otp.ExpiresAt)
	if err != nil {
		return nil, errors.New("OTP not found")
	}

	return otp, nil
}
