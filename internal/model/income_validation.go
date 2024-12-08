package model

func (i IncomeFormData) Validate() error {
	data := ModelEval{
		Floats: []Float{
			{
				Name:     "Amount",
				Value:    i.Amount,
				Positive: true,
			},
		},
		Strings: []String{
			{
				Name:  "Description",
				Value: i.Description,
				Min:   1,
				Max:   50,
			},
		},
		Ints: []Int{
			{
				Name:     "Entry class id",
				Value:    i.EntryClassID,
				Positive: true,
			},
		},
	}

	if err := data.Validate(); err != nil {
		return err
	}

	return nil
}
