-- +goose Up
-- +goose StatementBegin
CREATE TABLE reactors (
    id SERIAL PRIMARY KEY,
    farm_time INT NOT NULL,
    tokens_per_cycle INT NOT NULL,
    price INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reactors;
-- +goose StatementEnd
