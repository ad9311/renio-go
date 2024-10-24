package model

import eval "github.com/ad9311/renio-go/internal/eval"

var SignUpDataValidation = eval.ModelEval{
	String: eval.StringEval{
		"Username": {
			eval.Min: 4,
			eval.Max: 20,
		},
		"Name": {
			eval.Min: 4,
			eval.Max: 50,
		},
		"Email": {
			eval.Min: 7,
			eval.Max: 50,
		},
		"Password": {
			eval.Min: 8,
			eval.Max: 24,
		},
		"PasswordConfirmation": {
			eval.Min: 8,
			eval.Max: 24,
		},
	},
}

func (s SignUpData) Validate() error {
	if err := SignUpDataValidation.ValidateModel(s); err != nil {
		return err
	}

	return nil
}
