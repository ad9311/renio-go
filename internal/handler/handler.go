package handler

import (
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/justinas/nosurf"
)

type TmplData map[string]any

func writeTemplate(w http.ResponseWriter, name string, data TmplData) {
	cache := conf.GetTemplates()

	tmpl, ok := cache[name]
	if !ok {
		msg := fmt.Sprintf("template %s.tmpl.html not found", name)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Printf("RENDER %s.tmpl.html\n", name)
	err := tmpl.Execute(w, data)
	if err != nil {
		msg := fmt.Sprintf("error while rendering template, %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

func (td TmplData) SetCSRFToken(r *http.Request) {
	td["CSRFToken"] = nosurf.Token(r)
}
