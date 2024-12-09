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
	budget := GetBudgetCTX(ctx)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	GetAppDataCTX(ctx)["budget"] = budget
	GetAppDataCTX(ctx)["entryClasses"] = entryClasses
	writeTemplate(w, ctx, "expenses/new")
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := GetBudgetCTX(ctx)

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

	GetAppDataCTX(ctx)["budget"] = budget
	GetAppDataCTX(ctx)["entryClasses"] = entryClasses

	if _, err := svc.CreateExpense(expenseFormData, budget); err != nil {
		handleFormError(w, ctx, err, "expenses/new")
		return
	}

	GetAppDataCTX(ctx)["info"] = "Expense created successfully"
	w.WriteHeader(http.StatusCreated)
	writeTemplate(w, ctx, "expenses/new")
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := GetBudgetCTX(ctx)
	expense := GetExpenseCTX(ctx)

	GetAppDataCTX(ctx)["budget"] = budget
	GetAppDataCTX(ctx)["expense"] = expense
	GetAppDataCTX(ctx)["modalTitle"] = "Delete Expense"
	GetAppDataCTX(ctx)["confirmationMessage"] = "Are you sure you want to delete this expense?"
	GetAppDataCTX(ctx)["formAction"] = "/budgets/" + budget.UID + "/expenses/" + strconv.Itoa(expense.ID) + "/delete"
	writeTemplate(w, ctx, "expenses/show")
}

func GetEditExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := GetBudgetCTX(ctx)
	expense := GetExpenseCTX(ctx)

	var entryClasses model.EntryClasses
	if err := entryClasses.Index(); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	GetAppDataCTX(ctx)["budget"] = budget
	GetAppDataCTX(ctx)["entryClasses"] = entryClasses
	GetAppDataCTX(ctx)["expense"] = expense
	writeTemplate(w, ctx, "expenses/edit")
}

func PatchExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := GetBudgetCTX(ctx)
	expense := GetExpenseCTX(ctx)

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

	GetAppDataCTX(ctx)["budget"] = budget
	GetAppDataCTX(ctx)["entryClasses"] = entryClasses

	expense, err := svc.UpdateExpense(expense, expenseFormData, budget)
	GetAppDataCTX(ctx)["expense"] = expense
	if err != nil {
		handleFormError(w, ctx, err, "expenses/edit")
		return
	}

	GetAppDataCTX(ctx)["info"] = "Expense updated successfully"
	writeTemplate(w, ctx, "expenses/show")
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := GetBudgetCTX(ctx)
	expense := GetExpenseCTX(ctx)

	if err := svc.DeleteExpense(expense, budget); err != nil {
		writeInternalError(w, ctx, []string{err.Error()})
		return
	}

	http.Redirect(w, r, "/budgets/"+budget.UID, http.StatusSeeOther)
}
