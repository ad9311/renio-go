package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	data := TmplData{}
	data.SetCSRFToken(r)
	writeTemplate(w, "registration/index", data)
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

	_, err := svc.SignUpUser(signUnData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
}
