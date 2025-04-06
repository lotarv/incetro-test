
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance INT DEFAULT 0,
    active_reactor INT,
    farm_status VARCHAR(50) DEFAULT 'start',
    farm_start_time TIMESTAMP,
    farm_progress INT DEFAULT 0
);

CREATE TABLE reactors (
    id SERIAL PRIMARY KEY,
    farm_time INT NOT NULL,
    tokens_per_cycle INT NOT NULL,
    price INT NOT NULL
);

CREATE TABLE user_reactors (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    reactor_id INT REFERENCES reactors(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, reactor_id)
);

INSERT INTO reactors (farm_time, tokens_per_cycle, price) VALUES
(60, 10, 50), (120, 25, 100), (300, 70, 250);

INSERT INTO users (name, balance, active_reactor, farm_status) 
VALUES ('TestUser', 1000, 1, 'start');

INSERT INTO user_reactors (user_id, reactor_id) 
VALUES (1, 1);