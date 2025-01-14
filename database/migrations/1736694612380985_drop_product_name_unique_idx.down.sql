ALTER TABLE IF EXISTS food_products
ADD CONSTRAINT food_products_name_key UNIQUE (name);
