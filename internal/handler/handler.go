package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
)

type TmplData map[string]any

type NavLink struct {
	Name string
	URL  string
}

func GetNavLinks(ctx context.Context) []NavLink {
	return []NavLink{
		{Name: "Home", URL: "/home"},
		{Name: "Bugets", URL: "/budgets"},
	}
}

func SetSessionCTX(ctx context.Context, key vars.ContextKey, value any) {
	app.GetSession().Put(ctx, string(key), value)
}

func GetAppDataCTX(ctx context.Context) TmplData {
	return ctx.Value(vars.AppDataKey).(TmplData)
}

func GetUserCTX(ctx context.Context) model.SafeUser {
	session := app.GetSession()
	key := string(vars.CurrentUserKey)
	user, ok := session.Get(ctx, key).(model.SafeUser)
	if !ok {
		return model.SafeUser{}
	}
	return user
}

func IsUserSignedInCTX(ctx context.Context) bool {
	session := app.GetSession()
	key := string(vars.UserSignedInKey)
	return session.GetBool(ctx, key)
}

func GetUserID_CTX(ctx context.Context) int {
	session := app.GetSession()
	key := string(vars.UserIDKey)
	return session.GetInt(ctx, key)
}

func GetBudgetAccountCTX(ctx context.Context) model.BudgetAccount {
	return ctx.Value(vars.BudgetAccountKey).(model.BudgetAccount)
}

func GetBudgetCTX(ctx context.Context) model.Budget {
	return ctx.Value(vars.BudgetKey).(model.Budget)
}

func GetIncomeCTX(ctx context.Context) model.Income {
	return ctx.Value(vars.IncomeKey).(model.Income)
}

func GetExpenseCTX(ctx context.Context) model.Expense {
	return ctx.Value(vars.ExpenseKey).(model.Expense)
}

func writeTemplate(w http.ResponseWriter, ctx context.Context, name string) {
	cache := app.GetTemplates()
	tmpl, ok := cache[name]
	if !ok {
		msg := fmt.Sprintf("template %s.tmpl.html not found", name)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	data := GetAppDataCTX(ctx)
	executeTemplate(w, tmpl, name, data)
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

func writeInternalError(w http.ResponseWriter, ctx context.Context, errStrs []string) {
	w.WriteHeader(http.StatusInternalServerError)
	saveAppDataErrors(ctx, errStrs)
	writeTemplate(w, ctx, "error/index")
}

func writeAsBadRequest(w http.ResponseWriter, ctx context.Context, errStrs []string, page string) {
	w.WriteHeader(http.StatusBadRequest)
	saveAppDataErrors(ctx, errStrs)
	writeTemplate(w, ctx, page)
}

// --- Helpers --- //

func saveAppDataErrors(ctx context.Context, errStrs []string) {
	data := GetAppDataCTX(ctx)
	data["errors"] = errStrs
}

func handleFormError(w http.ResponseWriter, ctx context.Context, err error, tmpl string) {
	errEval, ok := err.(*model.ErrEval)
	if ok {
		GetAppDataCTX(ctx)["errors"] = errEval.Issues
		w.WriteHeader(http.StatusBadRequest)
		writeTemplate(w, ctx, tmpl)
		return
	} else {
		errStr := []string{err.Error()}
		writeInternalError(w, ctx, errStr)
		return
	}
}

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data TmplData) {
	fmt.Printf("RENDER %s.tmpl.html\n", name)

	err := tmpl.Execute(w, data)
	if err != nil {
		msg := fmt.Sprintf("error while rendering template, %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
