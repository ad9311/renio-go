package model

import (
	"github.com/ad9311/renio-go/internal/vars"
)

func (s *SignUpData) Validate() error {
	data := ModelEval{
		Strings: []String{
			{
				Name:  "Username",
				Value: s.Username,
				Min:   4,
				Max:   20,
			},
			{
				Name:  "Name",
				Value: s.Name,
				Min:   2,
				Max:   50,
			},
			{
				Name:    "Email",
				Value:   s.Email,
				Pattern: vars.EmailPattern,
			},
			{
				Name:  "Password",
				Value: s.Password,
				Min:   8,
				Max:   30,
			},
			{
				Name:  "Password confirmation",
				Value: s.PasswordConfirmation,
				Min:   8,
				Max:   30,
			},
		},
	}

	if err := data.Validate(); err != nil {
		return err
	}

	return nil
}
