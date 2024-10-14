package model

import (
	"context"
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type AllowedJWT struct {
	ID     int
	JTI    string
	AUD    string
	EXP    time.Time
	UserID int
}

type JWT struct {
	Token string
	JTI   string
	AUD   string
	EXP   time.Time
}

// --- Query --- //

func (aJWT *AllowedJWT) Insert(token JWT, userID int) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO allowed_jwts (jti, aud, exp, user_id)
						VALUES ($1, $2, $3, $4)`

	_, err := pool.Exec(ctx, query, token.JTI, token.AUD, token.EXP, userID)
	if err != nil {
		return err
	}

	return nil
}

func (aJWT *AllowedJWT) SelectByJTI(jit string) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `SELECT user_id FROM allowed_jwts WHERE jti = $1`

	err := pool.QueryRow(ctx, query, jit).Scan(&aJWT.UserID)
	if err != nil {
		return err
	}

	return nil
}
