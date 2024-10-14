package model

import (
	"context"
	"database/sql"

	"github.com/ad9311/renio-go/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int            `json:"id"`
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Image    sql.NullString `json:"image"`

	Password string
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

// Query functions //

func (u *User) Create(signUpData SignUpData) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO users (username, name, email, password)
						VALUES ($1, $2, $3, $4)
						RETURNING id, username, name, email, image`

	username := signUpData.Username
	name := signUpData.Name
	email := signUpData.Email
	password, err := hashPassword(signUpData.Password)
	if err != nil {
		return err
	}

	err = pool.QueryRow(ctx, query, username, name, email, password).Scan(
		&u.ID,
		&u.Username,
		&u.Name,
		&u.Email,
		&u.Image,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) FindForAuth(email string) error {
	query := `SELECT id, username, password FROM users WHERE email = $1`
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, email).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) SetUpAccounts() error {
	var budgetAccount BudgetAccount
	if err := budgetAccount.Insert(u.ID); err != nil {
		return err
	}

	return nil
}

func (u *User) FindByID(userID int) error {
	query := `SELECT id, username, name, email, image FROM users WHERE id = $1`
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, userID).Scan(
		&u.ID,
		&u.Username,
		&u.Name,
		&u.Email,
		&u.Image,
	)
	if err != nil {
		return err
	}

	return nil
}

// Helpers //

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
