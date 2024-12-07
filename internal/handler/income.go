package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func IncomeCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		incomeID := chi.URLParam(r, "incomeID")

		id, _ := strconv.Atoi(incomeID)
		var income model.Income

		err := income.SelectByID(id)
		if err == pgx.ErrNoRows {
			return
		}
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), vars.IncomeKey, income)
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

	if _, err := svc.CreateIncome(incomeFormData, budget); err != nil {
		errEval, ok := err.(*eval.ErrEval)
		if ok {
			getAppData(ctx)["errors"] = errEval.Issues
		} else {
			errStr := []string{err.Error()}
			writeInternalError(w, ctx, errStr)
			return
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		getAppData(ctx)["info"] = "Income created successfully"
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses
	writeTemplate(w, ctx, "income-list/new")
}
