-- +goose Up
-- +goose StatementBegin
ALTER TABLE entry_classes RENAME COLUMN "group" TO type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
