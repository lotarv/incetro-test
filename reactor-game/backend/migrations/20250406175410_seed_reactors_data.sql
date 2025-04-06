-- +goose Up
-- +goose StatementBegin
INSERT INTO reactors (farm_time, tokens_per_cycle, price) VALUES
(60, 10, 50),
(120, 25, 100),
(300, 70, 250);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM reactors WHERE id IN (1, 2, 3);
-- +goose StatementEnd
