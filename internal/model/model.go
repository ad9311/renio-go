package model

type ErrIncompleteQuery struct{}

func (e ErrIncompleteQuery) Error() string {
	return "query could not be finished"
}

// --- Validations --- //
