package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func IncomeCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		incomeID := chi.URLParam(r, "incomeID")

		id, _ := strconv.Atoi(incomeID)
		var income model.Income

		err := income.SelectByID(id)
		if err == pgx.ErrNoRows {
			writeNotFound(w, ctx)
			return
		}
		if err != nil {
			writeInternalError(w, ctx, []string{err.Error()})
			return
		}

		ctx = context.WithValue(ctx, vars.IncomeKey, income)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetNewIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses
	writeTemplate(w, ctx, "income-list/new")
}

func PostIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	if err := r.ParseForm(); err != nil {
		errStr := []string{err.Error()}
		writeAsBadRequest(w, ctx, errStr, "income-list/new")
		return
	}

	entryClasID, _ := strconv.Atoi(r.FormValue("entry_class_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 32)

	incomeFormData := model.IncomeFormData{
		EntryClassID: entryClasID,
		Description:  r.FormValue("description"),
		Amount:       float32(amount),
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses

	if _, err := svc.CreateIncome(incomeFormData, budget); err != nil {
		handleFormError(w, ctx, err, "income-list/new")
		return
	}

	getAppData(ctx)["info"] = "Income created successfully"
	w.WriteHeader(http.StatusCreated)
	writeTemplate(w, ctx, "income-list/new")
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	income := ctx.Value(vars.IncomeKey).(model.Income)

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["income"] = income
	writeTemplate(w, ctx, "income-list/show")
}

func GetEditIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	income := ctx.Value(vars.IncomeKey).(model.Income)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses
	getAppData(ctx)["income"] = income
	writeTemplate(w, ctx, "income-list/edit")
}

func PatchIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	income := ctx.Value(vars.IncomeKey).(model.Income)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	if err := r.ParseForm(); err != nil {
		errStr := []string{err.Error()}
		writeAsBadRequest(w, ctx, errStr, "income-list/edit")
		return
	}

	entryClassID, _ := strconv.Atoi(r.FormValue("entry_class_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 32)

	incomeFormData := model.IncomeFormData{
		EntryClassID: entryClassID,
		Description:  r.FormValue("description"),
		Amount:       float32(amount),
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses

	income, err := svc.UpdateIncome(income, incomeFormData, budget)
	getAppData(ctx)["income"] = income
	if err != nil {
		handleFormError(w, ctx, err, "income-list/edit")
		return
	}

	getAppData(ctx)["info"] = "Income updated successfully"
	writeTemplate(w, ctx, "income-list/show")
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	income := ctx.Value(vars.IncomeKey).(model.Income)

	if err := svc.DeleteIncome(income, budget); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	http.Redirect(w, r, "/budgets/"+budget.UID, http.StatusSeeOther)
}
