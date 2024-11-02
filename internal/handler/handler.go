package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/ad9311/renio-go/internal/conf"
)

func Root(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "sign-in.tmpl.html")
}

func writeTemplate(w http.ResponseWriter, name string) {
	cache := conf.GetTemplates(template.FuncMap{})

	tmpl, ok := cache[name]
	if !ok {
		fmt.Println("template does not exist")
	}

	err := tmpl.Execute(w, map[string]string{})
	if err != nil {
		fmt.Println(err)
	}
}
