CREATE TABLE IF NOT EXISTS dish_products
(
    dish_id         BIGSERIAL   NOT NULL REFERENCES dishes,
    food_product_id BIGSERIAL   NOT NULL REFERENCES food_products,
    amount          float4      NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_dish_products_on_dish_and_product ON dish_products (
    dish_id, food_product_id
);
