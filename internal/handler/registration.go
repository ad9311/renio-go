package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	writeTemplate(w, ctx, "registration/index")
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := r.ParseForm(); err != nil {
		writeErrorPage(w, ctx, []string{err.Error()})
		return
	}

	signUnData := model.SignUpData{
		Username:             r.FormValue("username"),
		Name:                 r.FormValue("name"),
		Email:                r.FormValue("email"),
		Password:             r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	_, err := svc.SignUpUser(signUnData)
	errEval, ok := err.(*model.ErrEval)
	if ok {
		writeAsBadRequest(w, ctx, errEval.Issues, "registration/index")
		return
	}

	if err != nil {
		errStr := []string{err.Error()}
		writeInternalError(w, ctx, errStr)
		return
	}

	http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
}
