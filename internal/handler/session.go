package handler

import "net/http"

func GetSignIn(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "sign-in.tmpl.html")
}
