package model

func (e ExpenseFormData) Validate() error {
	data := ModelEval{
		Floats: []Float{
			{
				Name:     "Amount",
				Value:    e.Amount,
				Positive: true,
			},
		},
		Strings: []String{
			{
				Name:  "Description",
				Value: e.Description,
				Min:   1,
				Max:   50,
			},
		},
		Ints: []Int{
			{
				Name:     "Entry class id",
				Value:    e.EntryClassID,
				Positive: true,
			},
		},
	}

	if err := data.Validate(); err != nil {
		return err
	}

	return nil
}
