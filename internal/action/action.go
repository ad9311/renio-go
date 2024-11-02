package action

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ad9311/renio-go/internal/cnsl"
)

type (
	Content map[string]any
	Errors  []string
)

type DataResponse struct {
	Content Content
	Status  int
}

type ErrorResponse struct {
	Errors Errors
	Status int
}

func WriteError(w http.ResponseWriter, errResponse ErrorResponse) {
	if errResponse.Status == 0 {
		errResponse.Status = 400
	}

	w.WriteHeader(errResponse.Status)

	data := Content{
		"errors": errResponse.Errors,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		cnsl.Error(fmt.Sprintf("could not write response, %s", err.Error()))
	}
}

func WriteOK(w http.ResponseWriter, dataResponse DataResponse) {
	if dataResponse.Status == 0 {
		dataResponse.Status = 200
	}

	w.WriteHeader(dataResponse.Status)

	data := Content{
		"data": dataResponse.Content,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		cnsl.Error(fmt.Sprintf("could not write response, %s", err.Error()))
	}
}

func DecodeJSON(body io.ReadCloser, data any) error {
	if err := json.NewDecoder(body).Decode(data); err != nil {
		return err
	}

	return nil
}

func (er *ErrorResponse) Append(err error) {
	er.Errors = append(er.Errors, err.Error())
}

func (er *ErrorResponse) AppendIssues(issues []string) {
	er.Errors = append(er.Errors, issues...)
}
