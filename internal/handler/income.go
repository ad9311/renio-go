package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ad9311/renio-go/internal/model"
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
