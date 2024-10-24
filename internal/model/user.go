package model

import (
	"database/sql"
	"time"

	"github.com/ad9311/renio-go/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string
	Image     sql.NullString `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
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

// --- Query Functions --- //

func (u *User) Insert(signUpData SignUpData) error {
	if err := signUpData.Validate(); err != nil {
		return err
	}

	query := "INSERT INTO users (username, name, email, password) VALUES ($1, $2, $3, $4) RETURNING *"

	password, err := hashPassword(signUpData.Password)
	if err != nil {
		return err
	}

	queryExec := db.QueryExe{
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

	queryExec := db.QueryExe{
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

	queryExec := db.QueryExe{
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

func (u *User) SetUpAccounts() error {
	var budgetAccount BudgetAccount
	if err := budgetAccount.Insert(u.ID); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) saveUserFromDB(queryExec db.QueryExe) error {
	value, ok := queryExec.Model.(*User)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*u = *value

	return nil
}
