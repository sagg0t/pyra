CREATE TABLE IF NOT EXISTS users
(
    id              BIGSERIAL           PRIMARY KEY,
    first_name      VARCHAR(64)         NOT NULL,
    last_name       VARCHAR(64)         NOT NULL,
    email           VARCHAR(128) UNIQUE NOT NULL,
    created_at      TIMESTAMPTZ         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ         NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_on_email
ON users (email);
