CREATE TABLE IF NOT EXISTS food_products
(
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(32) UNIQUE NOT NULL,
    proteins        float8             NOT NULL DEFAULT 0,
    fats            float8             NOT NULL DEFAULT 0,
    carbos          float8             NOT NULL DEFAULT 0,
    calories        float8             NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP
);
