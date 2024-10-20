-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS allowed_jwts (
    id SERIAL PRIMARY KEY,
    jti VARCHAR(255) NOT NULL,
    aud VARCHAR(255) NOT NULL,
    exp TIMESTAMPTZ NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
      FOREIGN KEY (user_id)
      REFERENCES users(id)
      ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS allowed_jwts;
-- +goose StatementEnd
