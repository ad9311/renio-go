package model

import (
	"database/sql"
	"time"
)

// User models

type User struct {
	ID       int            `json:"id"`
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Image    sql.NullString `json:"image"`
}

type SignUpData struct {
	Username             string `json:"username"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Allowed JWT

type AllowedJWT struct {
	ID     int
	JTI    string
	AUD    string
	EXP    time.Time
	UserID int
}

type NewJWT struct {
	Token string
	JTI   string
	AUD   string
	EXP   time.Time
}
