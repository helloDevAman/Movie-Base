package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/helloDevAman/movie-base/internal/domain/entities"
)

type OTPRepository interface {
	SaveOTP(otp entities.OTP) error
	GetLatestOTP(mobileNumber string) (*entities.OTP, error)
}

type PostgresOTPRepository struct {
	db *sql.DB
}

func NewPostgresOTPRepository(db *sql.DB) *PostgresOTPRepository {
	return &PostgresOTPRepository{db: db}
}

func (r *PostgresOTPRepository) SaveOTP(otp entities.OTP) error {
	query := `
		INSERT INTO otps (mobile, code, created_at, expires_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (mobile) DO UPDATE 
		SET code = EXCLUDED.code, created_at = EXCLUDED.created_at, expires_at = EXCLUDED.expires_at;
	`
	_, err := r.db.ExecContext(context.Background(), query, otp.MobileNumber, otp.Code, otp.CreatedAt, otp.ExpiresAt)
	return err
}

func (r *PostgresOTPRepository) GetLatestOTP(mobileNumber string) (*entities.OTP, error) {
	query := `SELECT mobile, code, created_at, expires_at FROM otps WHERE mobile = $1`
	row := r.db.QueryRowContext(context.Background(), query, mobileNumber)

	var otp entities.OTP
	err := row.Scan(&otp.MobileNumber, &otp.Code, &otp.CreatedAt, &otp.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("OTP not found")
		}
		return nil, err
	}

	// Check if OTP has expired
	if time.Now().After(otp.ExpiresAt) {
		return nil, errors.New("OTP has expired")
	}

	return &otp, nil
}
