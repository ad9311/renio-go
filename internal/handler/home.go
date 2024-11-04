package handler

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, r, "home/index")
}
