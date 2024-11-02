package action

import (
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/jackc/pgx/v5"
)

// --- Actions --- //

func IndexBudgets(w http.ResponseWriter, r *http.Request) {
	budgetAccount := r.Context().Value(vars.BudgetAccountKey).(model.BudgetAccount)

	errResponse := ErrorResponse{}
	budgets := model.Budgets{}
	if err := budgets.Index(budgetAccount.ID); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"budgets": budgets,
		},
	}

	WriteOK(w, dataResponse)
}

func PostBudget(w http.ResponseWriter, r *http.Request) {
	budgetAccount := r.Context().Value(vars.BudgetAccountKey).(model.BudgetAccount)

	errResponse := ErrorResponse{}
	var budget model.Budget
	if err := budget.Insert(budgetAccount.ID); err != nil {
		err = fmt.Errorf("budget already exists")
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"budget": budget,
		},
		Status: http.StatusCreated,
	}

	WriteOK(w, dataResponse)
}

func GetBudget(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	query := r.URL.Query().Get("entries")
	options := svc.ParseBudgetEntriesQuery(query)

	errResponse := ErrorResponse{}
	budgetWithEntries, err := svc.FindBudgetWithEntries(budget, options)
	if err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"budget": budgetWithEntries,
		},
	}

	WriteOK(w, dataResponse)
}

func GetCurrentBudget(w http.ResponseWriter, r *http.Request) {
	budgetAccount := r.Context().Value(vars.BudgetAccountKey).(model.BudgetAccount)

	errResponse := ErrorResponse{}
	var budget model.Budget
	err := budget.SelectCurrent(budgetAccount.ID)
	if err == pgx.ErrNoRows {
		err = fmt.Errorf("budget not found")
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"budget": budget,
		},
	}

	WriteOK(w, dataResponse)
}
