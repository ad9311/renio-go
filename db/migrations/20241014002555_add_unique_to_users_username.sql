-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD CONSTRAINT users_unsername_key UNIQUE (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP CONSTRAINT users_unsername_key;
-- +goose StatementEnd
