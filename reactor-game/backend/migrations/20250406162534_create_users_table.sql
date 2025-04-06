-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance INT DEFAULT 0,
    active_reactor INT,
    farm_status VARCHAR(50) DEFAULT 'start',
    farm_start_time TIMESTAMP,
    farm_progress INT DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
