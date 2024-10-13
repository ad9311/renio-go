package usermodel

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/lib"
	"github.com/ad9311/renio-go/internal/model"
)

type UserToConfirm struct {
	ID             int
	Username       string
	HashedPassword string
}

func Create(signUpData model.SignUpData) (model.User, error) {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO users (username, name, email, password)
						VALUES ($1, $2, $3, $4)
						RETURNING id, username, name, email, image`

	var user model.User

	username := signUpData.Username
	name := signUpData.Name
	email := signUpData.Email
	password, err := lib.HashPassword(signUpData.Password)
	if err != nil {
		return user, err
	}

	err = pool.QueryRow(ctx, query, username, name, email, password).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Image,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func FindForAuth(email string) (UserToConfirm, error) {
	query := `SELECT id, username, password FROM users WHERE email = $1`
	var user UserToConfirm
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.HashedPassword)

	if err != nil {
		return user, err
	}

	return user, nil
}
