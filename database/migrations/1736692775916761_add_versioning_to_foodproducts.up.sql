ALTER TABLE IF EXISTS food_products
ADD COLUMN IF NOT EXISTS uid UUID NOT NULL DEFAULT gen_random_uuid(),
ADD COLUMN IF NOT EXISTS version int4 NOT NULL DEFAULT 1;

ALTER TABLE IF EXISTS food_products
ALTER COLUMN version DROP DEFAULT,
ALTER COLUMN uid DROP DEFAULT;
