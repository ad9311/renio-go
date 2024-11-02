package handler

import "net/http"

func GetHome(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "home.tmpl.html")
}
