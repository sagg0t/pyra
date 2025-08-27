package products

const productColumns = "id, uid, version, name, calories, proteins, fats, carbs, created_at, updated_at"

const indexProductsQuery = "SELECT " + productColumns + " FROM food_products;"

const findByIDQuery = "SELECT " + productColumns + " FROM food_products WHERE id = $1 LIMIT 1"

const findByRefQuery = "SELECT " + productColumns + " FROM food_products WHERE uid = $1 AND version = $2 LIMIT 1"

const productVersionsQuery = "SELECT " + productColumns + " FROM food_products WHERE uid = $1"

const findAllByIDsQuery = "SELECT " + productColumns + " FROM food_products WHERE id in $0"

const productsForDishQuery = "SELECT " + productColumns + ` FROM dish_products
JOIN food_products ON dish_products.food_product_id = food_products.id
WHERE dish_products.dish_id = $1;`

const createProductVersionQuery = `INSERT INTO food_products (
	uid, version, name, calories, proteins, fats, carbs
) VALUES (
	$1, (SELECT MAX(version) + 1 FROM food_products WHERE uid = $1),
	$2, $3, $4, $5, $6
) RETURNING id, version, created_at;`

const createProductQuery = `INSERT INTO food_products (
	uid, version, name, calories, proteins, fats, carbs
) VALUES (
	$1, 1, $2, $3, $4, $5, $6
) RETURNING id, version;`

const deleteByIDQuery = "DELETE FROM food_products WHERE id = $1"

const updateProductQuery = `UPDATE food_products
SET name = $2, calories = $3, proteins = $4,
	fats = $5, carbs = $6
WHERE id = $1
RETURNING uid, version, created_at, updated_at;`

const searchProductsQuery = "SELECT " + productColumns + ` FROM food_products
WHERE name ILIKE '%' || $1 || '%'`

const nameTakenQuery = "SELECT 1 FROM food_products WHERE name LIKE $1"

const maxProductVersionQuery = "SELECT COALESCE(MAX(version), 0) FROM food_products WHERE uid = $1"
