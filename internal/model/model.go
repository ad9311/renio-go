package model

import (
	"time"
)

// Allowed JWT

type AllowedJWT struct {
	ID     int
	JTI    string
	AUD    string
	EXP    time.Time
	UserID int
}
