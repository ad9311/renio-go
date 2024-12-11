package model

import (
	"database/sql"
	"time"

	"github.com/ad9311/renio-go/internal/app"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Username  string
	Name      string
	Email     string
	Password  string
	Image     sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SafeUser struct {
	Username string
	Name     string
	Email    string
	Image    sql.NullString
}

type SignUpData struct {
	Username             string
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

type SignInData struct {
	Email    string
	Password string
}

// --- Query Functions --- //

func (u *User) Insert(signUpData SignUpData) error {
	query := `
	INSERT INTO users (username, name, email, password)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	password, err := hashPassword(signUpData.Password)
	if err != nil {
		return err
	}

	queryExec := app.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			signUpData.Username,
			signUpData.Name,
			signUpData.Email,
			password,
		},
		Model: User{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := u.saveUserFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (u *User) SelectByEmail(email string) error {
	query := "SELECT * FROM users WHERE email = $1"

	queryExec := app.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{email},
		Model:     User{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := u.saveUserFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (u *User) SelectByID(userID int) error {
	query := "SELECT * FROM users WHERE id = $1"

	queryExec := app.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{userID},
		Model:     User{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := u.saveUserFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (u *User) GetSafeUser() SafeUser {
	return SafeUser{
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
		Image:    u.Image,
	}
}

// --- Helpers --- //

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) saveUserFromDB(queryExec app.QueryExe) error {
	value, ok := queryExec.Model.(*User)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*u = *value

	return nil
}
