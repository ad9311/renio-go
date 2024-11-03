package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/ad9311/renio-go/internal/conf"
)

func writeTemplate(w http.ResponseWriter, name string) {
	cache := conf.GetTemplates(template.FuncMap{})

	tmpl, ok := cache[name]
	if !ok {
		msg := fmt.Sprintf("template %s.tmpl.html not found", name)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Printf("RENDER %s.tmpl.html\n", name)
	err := tmpl.Execute(w, map[string]string{})
	if err != nil {
		msg := fmt.Sprintf("error while rendering template, %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
