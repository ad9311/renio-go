package usermodel

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

func FindUserByCreds(email string, password string) (*User, error) {
	var user User
	pool := db.GetPool()
	query := `SELECT id, username, email, image FROM users WHERE email = $1`
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Image)
	if err != nil {
		return nil, err
	}

	return &user, err
}
