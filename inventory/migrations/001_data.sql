-- +goose up
CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL,
    status VARCHAR(30) NOT NULL,
    number INT NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO inventory (status, number) VALUES ('active', 10);

-- +goose down
DROP TABLE inventory;