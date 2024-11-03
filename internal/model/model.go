package model

import "encoding/gob"

type ErrIncompleteQuery struct{}

func (e ErrIncompleteQuery) Error() string {
	return "query could not be finished"
}

func RegisterModels() {
	gob.Register(SafeUser{})
}
