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

	errResponse := ErrorResponse{}
	if err := DecodeJSON(r.Body, &signUpData); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	err := svc.SignUpUser(signUpData)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		errResponse.AppendIssues(errEval.Issues)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"message": "User created successfully",
		},
		Status: http.StatusCreated,
	}
	WriteOK(w, dataResponse)
}
