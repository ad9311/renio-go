package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
)

func GetSignIn(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, r, "session/index")
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	signInData := model.SignInData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	user, err := svc.SignInUser(signInData)
	if err != nil {
		GetAppData(r)["errors"] = []string{err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		writeTemplate(w, r, "session/index")
		return
	}

	conf.GetSession().Put(r.Context(), string(vars.UserSignedInKey), true)
	conf.GetSession().Put(r.Context(), string(vars.CurrentUserKey), user)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func PostSignOut(w http.ResponseWriter, r *http.Request) {
	_ = conf.GetSession().Destroy(r.Context())
	_ = conf.GetSession().RenewToken(r.Context())
	http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
}
