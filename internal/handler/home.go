package handler

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	data := TmplData{}
	data.SetCurrentUser(r)
	writeTemplate(w, "home/index", data)
}
