-- +goose up
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL,
    status VARCHAR(30) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +goose down
DROP TABLE orders;