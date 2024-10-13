package usermodel

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/lib"
	"github.com/ad9311/renio-go/internal/model"
)

func Create(signUpData model.SignUpData) (*model.User, error) {
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
		return nil, err
	}

	err = pool.QueryRow(ctx, query, username, name, email, password).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Image,
	)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func FindUserByCreds(email string, password string) (*model.User, error) {
	query := `SELECT id, username, email, image FROM users WHERE email = $1`
	return findUser(query, email)
}

func FindUserByID(id int) (*model.User, error) {
	query := `SELECT id, username, email, image FROM users WHERE id = $1`
	return findUser(query, id)
}

func findUser(query string, arg any) (*model.User, error) {
	var user model.User
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, arg).Scan(&user.ID, &user.Username, &user.Email, &user.Image)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
