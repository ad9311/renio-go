package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type AllowedJWT struct {
	ID        int
	JTI       string
	AUD       string
	EXP       time.Time
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JWT struct {
	Token string
	JTI   string
	AUD   string
	EXP   time.Time
}

// --- Query --- //

func (a *AllowedJWT) Insert(token JWT, userID int) error {
	query := "INSERT INTO allowed_jwts (jti, aud, exp, user_id) VALUES ($1, $2, $3, $4) RETURNING *"

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			token.JTI,
			token.AUD,
			token.EXP,
			userID,
		},
		Model: AllowedJWT{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := a.saveAllowedJWTFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (a *AllowedJWT) SelectByJTI(jti string) error {
	query := `SELECT * FROM allowed_jwts WHERE jti = $1`

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{jti},
		Model:     AllowedJWT{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := a.saveAllowedJWTFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (a *AllowedJWT) saveAllowedJWTFromDB(queryExec db.QueryExe) error {
	value, ok := queryExec.Model.(*AllowedJWT)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*a = *value

	return nil
}
