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

func (aJWT *AllowedJWT) Insert() error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO allowed_jwts (jti, aud, exp, user_id)
						VALUES ($1, $2, $3, $4)`

	_, err := pool.Exec(ctx, query, aJWT.JTI, aJWT.AUD, aJWT.EXP, aJWT.UserID)
	if err != nil {
		return err
	}

	return nil
}
