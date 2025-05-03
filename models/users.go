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
