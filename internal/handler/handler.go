package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/vars"
)

type TmplData map[string]any

func writeTemplate(w http.ResponseWriter, r *http.Request, name string) {
	cache := conf.GetTemplates()
	tmpl, ok := cache[name]
	if !ok {
		msg := fmt.Sprintf("template %s.tmpl.html not found", name)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	data := GetAppData(r.Context())
	executeTemplate(w, tmpl, name, data)
}

func GetAppData(ctx context.Context) TmplData {
	return ctx.Value(vars.AppDataKey).(TmplData)
}

func GetCurrentUserId(ctx context.Context) int {
	userIDkey := string(vars.UserIDKey)
	return conf.GetSession().GetInt(ctx, userIDkey)
}

// --- Helpers --- //

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data TmplData) {
	fmt.Printf("RENDER %s.tmpl.html\n", name)

	err := tmpl.Execute(w, data)
	if err != nil {
		msg := fmt.Sprintf("error while rendering template, %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
