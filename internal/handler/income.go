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

func GetIncomeForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		GetAppData(ctx).AppendError(ctx, err)
		writeTemplate(w, r, "error/index")
		return
	}

	GetAppData(ctx)["budget"] = budget
	GetAppData(ctx)["entryClasses"] = entryClasses
	writeTemplate(w, r, "income-list/new")
}

func PostIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	if err := r.ParseForm(); err != nil {
		GetAppData(ctx).AppendError(ctx, err)
		w.WriteHeader(http.StatusBadRequest)
		writeTemplate(w, r, "error/index")
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
			GetAppData(ctx)["errors"] = errEval.Issues
		} else {
			GetAppData(ctx).AppendError(ctx, err)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	writeTemplate(w, r, "income-list/new")
}
