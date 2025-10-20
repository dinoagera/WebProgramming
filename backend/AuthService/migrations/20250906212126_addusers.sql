-- -- +goose Up
CREATE TABLE IF NOT EXISTS users (
    uid SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    pass_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS goose_db_version (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN NOT NULL,
    tstamp TIMESTAMP DEFAULT NOW()
);
-- -- +goose Down
DROP TABLE users;
