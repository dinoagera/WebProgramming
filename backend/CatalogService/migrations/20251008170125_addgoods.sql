-- +goose Up
CREATE TABLE IF NOT EXISTS goods (
    product_id SERIAL PRIMARY KEY,
    category VARCHAR(50) NOT NULL,
    sex VARCHAR(50) NOT NULL CHECK (sex IN ('male', 'female', 'unisex')),
    sizes INTEGER[] NOT NULL,
    price NUMERIC NOT NULL,
    color VARCHAR(50) NOT NULL,
    tag VARCHAR(50) NOT NULL,
    imageData bytea NOT NULL
);
CREATE TABLE IF NOT EXISTS goose_db_version (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN NOT NULL,
    tstamp TIMESTAMP DEFAULT NOW()
);
-- +goose Down
DROP TABLE goods;