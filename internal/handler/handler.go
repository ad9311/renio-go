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

func writeTemplate(w http.ResponseWriter, ctx context.Context, name string) {
	cache := conf.GetTemplates()
	tmpl, ok := cache[name]
	if !ok {
		msg := fmt.Sprintf("template %s.tmpl.html not found", name)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	data := getAppData(ctx)
	executeTemplate(w, tmpl, name, data)
}

func writeTurboTemplate(w http.ResponseWriter, ctx context.Context, name string, status int) {
	data := getAppData(ctx)
	data["turboTemplate"] = true
	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html; charset=utf-8")
	w.WriteHeader(status)
	writeTemplate(w, ctx, name)
}

// --- Helpers --- //

func getAppData(ctx context.Context) TmplData {
	return ctx.Value(vars.AppDataKey).(TmplData)
}

func getCurrentUserId(ctx context.Context) int {
	userIDkey := string(vars.UserIDKey)
	return conf.GetSession().GetInt(ctx, userIDkey)
}

func saveAppDataErrors(ctx context.Context, errors []string) {
	data := getAppData(ctx)
	data["errors"] = errors
}

func writeErrorPage(w http.ResponseWriter, ctx context.Context, errors []string) {
	w.WriteHeader(http.StatusBadRequest)
	saveAppDataErrors(ctx, errors)
	writeTemplate(w, ctx, "error/index")
}

func writeNotFound(w http.ResponseWriter, ctx context.Context) {
	w.WriteHeader(http.StatusNotFound)
	saveAppDataErrors(ctx, []string{"could not find page"})
	writeTemplate(w, ctx, "not-found/index")
}

func writeInternalError(w http.ResponseWriter, ctx context.Context, errors []string) {
	w.WriteHeader(http.StatusInternalServerError)
	saveAppDataErrors(ctx, errors)
	writeTemplate(w, ctx, "error/index")
}

func writeAsBadRequest(w http.ResponseWriter, ctx context.Context, errors []string, page string) {
	w.WriteHeader(http.StatusBadRequest)
	saveAppDataErrors(ctx, errors)
	writeTemplate(w, ctx, page)
}

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data TmplData) {
	fmt.Printf("RENDER %s.tmpl.html\n", name)

	err := tmpl.Execute(w, data)
	if err != nil {
		msg := fmt.Sprintf("error while rendering template, %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
