package utils

import (
	"database/sql"
	"fmt"
)

func VerifyOTP(db *sql.DB, email string, otp string) error {
	var dbOtp string
	err := db.QueryRow("SELECT otp FROM otp WHERE email=$1", email).Scan(&dbOtp)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if otp != dbOtp {
		return fmt.Errorf("invalid otp")
	}

	_, _ = db.Exec("DELETE FROM otp WHERE email=$1", email)
	return nil
}
