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

func ExpenseCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expenseID := chi.URLParam(r, "expenseID")

		id, _ := strconv.Atoi(expenseID)
		var expense model.Expense

		err := expense.SelectByID(id)
		if err == pgx.ErrNoRows {
			return
		}
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), vars.ExpenseKey, expense)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetNewExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses
	writeTemplate(w, ctx, "expenses/new")
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	if err := r.ParseForm(); err != nil {
		errStr := []string{err.Error()}
		writeAsBadRequest(w, ctx, errStr, "expenses/new")
		return
	}

	entryClasID, _ := strconv.Atoi(r.FormValue("entry_class_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 32)

	expenseFormData := model.ExpenseFormData{
		EntryClassID: entryClasID,
		Description:  r.FormValue("description"),
		Amount:       float32(amount),
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses

	if _, err := svc.CreateExpense(expenseFormData, budget); err != nil {
		handleFormError(w, ctx, err, "expenses/new")
		return
	}

	getAppData(ctx)["info"] = "Expense created successfully"
	w.WriteHeader(http.StatusCreated)
	writeTemplate(w, ctx, "expenses/new")
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	expense := ctx.Value(vars.ExpenseKey).(model.Expense)

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["expense"] = expense
	writeTemplate(w, ctx, "expenses/show")
}

func GetEditExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	expense := ctx.Value(vars.ExpenseKey).(model.Expense)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses
	getAppData(ctx)["expense"] = expense
	writeTemplate(w, ctx, "expenses/edit")
}

func PatchExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	expense := ctx.Value(vars.ExpenseKey).(model.Expense)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	if err := r.ParseForm(); err != nil {
		errStr := []string{err.Error()}
		writeAsBadRequest(w, ctx, errStr, "expenses/edit")
		return
	}

	entryClassID, _ := strconv.Atoi(r.FormValue("entry_class_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 32)

	expenseFormData := model.ExpenseFormData{
		EntryClassID: entryClassID,
		Description:  r.FormValue("description"),
		Amount:       float32(amount),
	}

	getAppData(ctx)["budget"] = budget
	getAppData(ctx)["entryClasses"] = entryClasses

	expense, err := svc.UpdateExpense(expense, expenseFormData, budget)
	getAppData(ctx)["expense"] = expense
	if err != nil {
		handleFormError(w, ctx, err, "expenses/edit")
		return
	}

	getAppData(ctx)["info"] = "Expense updated successfully"
	writeTemplate(w, ctx, "expenses/show")
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)
	expense := ctx.Value(vars.ExpenseKey).(model.Expense)

	if err := svc.DeleteExpense(expense, budget); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	http.Redirect(w, r, "/budgets/"+budget.UID, http.StatusSeeOther)
}
