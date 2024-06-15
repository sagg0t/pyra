CREATE TABLE IF NOT EXISTS ingredients
(
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR(32) UNIQUE NOT NULL,
    proteins      float8             NOT NULL DEFAULT 0,
    fats          float8             NOT NULL DEFAULT 0,
    carbohydrates float8             NOT NULL DEFAULT 0,
    calories      float8             NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP
);
