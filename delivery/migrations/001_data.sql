-- +goose up
CREATE TABLE IF NOT EXISTS delivery (
    id SERIAL,
    id_order INT NOT NULL,
    id_payment INT NOT NULL,
    status VARCHAR(30) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +goose down
DROP TABLE delivery;