package action

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

// --- Actions --- //

func PostUser(w http.ResponseWriter, r *http.Request) {
	var signUpData model.SignUpData
	if err := DecodeJSON(r.Body, &signUpData); err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	err := svc.SignUpUser(signUpData)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		WriteError(w, errEval.Issues, http.StatusBadRequest)
		return
	}
	if err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	WriteOK(w, "User created successfully", http.StatusCreated)
}
