package model

import (
	"github.com/ad9311/renio-go/internal/eval"
)

func (i IncomeFormData) Validate() eval.Issues {
	data := eval.ModelEval{
		Floats: []eval.Float{
			{
				Name:     "Amount",
				Value:    i.Amount,
				Positive: true,
			},
		},
		Strings: []eval.String{
			{
				Name:  "Description",
				Value: i.Description,
				Min:   1,
				Max:   50,
			},
		},
		Ints: []eval.Int{
			{
				Name:     "Entry class id",
				Value:    i.EntryClassID,
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
