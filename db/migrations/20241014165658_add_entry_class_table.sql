-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS entry_classes (
    id SERIAL PRIMARY KEY,
    uid VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    "group" INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transaction_types;
-- +goose StatementEnd
