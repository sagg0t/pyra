CREATE TABLE IF NOT EXISTS food_products
(
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(64) UNIQUE NOT NULL,
    calories        float4             NOT NULL DEFAULT 0,
    proteins        float4             NOT NULL DEFAULT 0,
    fats            float4             NOT NULL DEFAULT 0,
    carbs           float4             NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP
);
