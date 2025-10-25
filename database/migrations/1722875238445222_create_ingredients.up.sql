CREATE TABLE IF NOT EXISTS ingredients
(
    dish_id             BIGSERIAL   NOT NULL REFERENCES dishes,

    ingredientable_id   BIGSERIAL   NOT NULL,
    ingredientable_type int2        NOT NULL,

    amount              int4        NOT NULL,
    unit                int2        NOT NULL,

    idx                 int2        NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_ingredients_on_dish_and_ingredientable ON ingredients (
    dish_id, ingredientable_id, ingredientable_type
);
