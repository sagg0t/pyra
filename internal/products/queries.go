package products

const productColumns = "id, uid, version, name, calories, proteins, fats, carbs, created_at, updated_at"

const indexProductsQuery = `SELECT 
		products.id,
		products.uid,
		products.version,
		products.name,
		products.calories,
		products.proteins,
		products.fats,
		products.carbs,
		products.created_at,
		products.updated_at
	FROM products
	INNER JOIN (
		SELECT DISTINCT uid, max(version) AS version
		FROM products
		GROUP BY uid
	) latest_products ON products.uid = latest_products.uid
	AND products.version = latest_products.version
	ORDER BY created_at DESC;`

const findByIDQuery = "SELECT " + productColumns + " FROM products WHERE id = $1 LIMIT 1;"

const findByRefQuery = "SELECT " + productColumns + ` FROM products
WHERE uid = $1 AND version = $2
LIMIT 1;`

const productVersionsQuery = "SELECT " + productColumns + ` FROM products
	WHERE uid = $1
	ORDER BY version DESC
	LIMIT 20;`

const findAllByIDsQuery = "SELECT " + productColumns + " FROM products WHERE id in $0;"

const productsForDishQuery = "SELECT " + productColumns + ` FROM ingredients
JOIN products
ON ingredients.ingredientable_id = products.id
	AND ingredients.ingredientable_type = 1
WHERE ingredients.dish_id = $1;`

const createProductVersionQuery = `
INSERT INTO products (
	uid, version, name, calories, proteins, fats, carbs
) VALUES (
	$1, (SELECT MAX(version) + 1 FROM products WHERE uid = $1),
	$2, $3, $4, $5, $6
) RETURNING id, version, created_at;`

const createProductQuery = `INSERT INTO products (
	uid, version, name, calories, proteins, fats, carbs
) VALUES (
	$1, 1, $2, $3, $4, $5, $6
) RETURNING id, version;`

const deleteByIDQuery = "DELETE FROM products WHERE id = $1"

const updateProductQuery = `UPDATE products
SET name = $3, calories = $4, proteins = $5,
	fats = $6, carbs = $7
WHERE uid = $1 AND version = $2
RETURNING uid, version, created_at, updated_at;`

const searchProductsQuery = "SELECT " + productColumns + ` FROM products
WHERE name ILIKE '%' || $1 || '%'`

const nameTakenQuery = "SELECT 1 FROM products WHERE name LIKE $1"

const maxProductVersionQuery = "SELECT COALESCE(MAX(version), 0) FROM products WHERE uid = $1"

const usedInDishesQuery = `SELECT 1
FROM ingredients
WHERE ingredientable_id = $1 
	AND ingredientable_type = $2
LIMIT 1;`
