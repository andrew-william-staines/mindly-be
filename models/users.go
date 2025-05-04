package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      sql.NullString    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Password struct {
	Email     string    `json:"email"`
	OTP 	  string 	`json:"otp"`
	Password  string 	`json:"newPassword"`
}
