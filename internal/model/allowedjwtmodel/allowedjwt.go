package allowedjwtmodel

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/model"
)

func Create(allowedJWT model.AllowedJWT) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO allowed_jwts (jti, aud, exp, user_id)
						VALUES ($1, $2, $3, $4)`

	_, err := pool.Exec(ctx, query, allowedJWT.JTI, allowedJWT.AUD, allowedJWT.EXP, allowedJWT.UserID)
	if err != nil {
		return err
	}

	return nil
}
