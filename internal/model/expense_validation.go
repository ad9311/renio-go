package model

import (
	"github.com/ad9311/renio-go/internal/eval"
)

func (e ExpenseFormData) Validate() eval.Issues {
	data := eval.ModelEval{
		Floats: []eval.Float{
			{
				Name:     "Amount",
				Value:    e.Amount,
				Positive: true,
			},
		},
		Strings: []eval.String{
			{
				Name:  "Description",
				Value: e.Description,
				Min:   1,
				Max:   50,
			},
		},
		Ints: []eval.Int{
			{
				Name:     "Entry class id",
				Value:    e.EntryClassID,
				Positive: true,
			},
		},
	}

	issues := data.Validate()
	if issues != nil {
		return issues
	}

	return nil
}
