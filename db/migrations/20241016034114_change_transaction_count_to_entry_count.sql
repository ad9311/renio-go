-- +goose Up
-- +goose StatementBegin
ALTER TABLE budgets RENAME COLUMN transaction_count TO entry_count;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
