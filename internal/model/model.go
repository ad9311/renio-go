package model

import "database/sql"

type User struct {
	ID       int64          `json:"id"`
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Image    sql.NullString `json:"image"`
}

type SignUpData struct {
	Username             string `json:"username"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
