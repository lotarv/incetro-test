-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN telegram_id BIGINT UNIQUE;
UPDATE users SET telegram_id = 123456789 WHERE id = 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN telegram_id;
-- +goose StatementEnd
