-- +goose Up
-- +goose StatementBegin
ALTER TABLE allowed_jwts
ADD CONSTRAINT allowed_jwts_jti_key UNIQUE (jti);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE allowed_jwts
DROP CONSTRAINT allowed_jwts_jti_key;
-- +goose StatementEnd
