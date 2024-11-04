package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, r, "registration/index")
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signUnData := model.SignUpData{
		Username:             r.FormValue("username"),
		Name:                 r.FormValue("name"),
		Email:                r.FormValue("email"),
		Password:             r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	ctx := r.Context()
	_, err := svc.SignUpUser(signUnData)
	if err != nil {
		errEval, ok := err.(*eval.ErrEval)
		if ok {
			GetAppData(ctx)["errors"] = errEval.Issues
		} else {
			GetAppData(ctx)["errors"] = []string{err.Error()}
		}
		w.WriteHeader(http.StatusBadRequest)
		writeTemplate(w, r, "registration/index")
		return
	}

	http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
}
