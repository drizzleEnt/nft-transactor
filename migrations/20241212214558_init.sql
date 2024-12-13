-- +goose Up
CREATE TABLE tokens(
    id SERIAL PRIMARY KEY,
    unique_hash VARCHAR(256),
    tx_hash VARCHAR(256),
    media_url VARCHAR(256),
    owner VARCHAR(256)
);
-- +goose Down
DROP TABLE tokens;