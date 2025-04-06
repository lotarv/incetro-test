-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_reactors (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    reactor_id INT REFERENCES reactors(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, reactor_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_reactors;
-- +goose StatementEnd
