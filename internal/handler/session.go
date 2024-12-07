package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
)

func GetSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	writeTemplate(w, ctx, "session/index")
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		writeErrorPage(w, ctx, []string{err.Error()})
		return
	}

	signInData := model.SignInData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	user, err := svc.SignInUser(signInData)
	if err != nil {
		errStr := []string{"invalid email or password"}
		writeAsBadRequest(w, ctx, errStr, "session/index")
		return
	}

	conf.GetSession().Put(ctx, string(vars.UserSignedInKey), true)
	conf.GetSession().Put(ctx, string(vars.CurrentUserKey), user.GetSafeUser())
	conf.GetSession().Put(ctx, string(vars.UserIDKey), user.ID)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func PostSignOut(w http.ResponseWriter, r *http.Request) {
	_ = conf.GetSession().Destroy(r.Context())
	_ = conf.GetSession().RenewToken(r.Context())
	http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
}
