-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, balance, active_reactor, farm_status) 
VALUES ('TestUser', 1000, 1, 'start');

INSERT INTO user_reactors (user_id, reactor_id) 
VALUES (1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_reactors WHERE user_id = 1 AND reactor_id = 1;
DELETE FROM users WHERE id = 1;
-- +goose StatementEnd
