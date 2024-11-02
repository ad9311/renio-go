-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS budgets (
    id SERIAL PRIMARY KEY,
    uid VARCHAR(255) NOT NULL UNIQUE,
    balance NUMERIC(11,2) NOT NULL DEFAULT 0,
    total_income NUMERIC(11,2) NOT NULL DEFAULT 0,
    total_expenses NUMERIC(11,2) NOT NULL DEFAULT 0,
    transaction_count INT NOT NULL DEFAULT 0,
    income_count INT NOT NULL DEFAULT 0,
    expense_count INT NOT NULL DEFAULT 0,
    budget_account_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_budget_account
      FOREIGN KEY (budget_account_id)
      REFERENCES budget_accounts(id)
      ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS budgets;
-- +goose StatementEnd
