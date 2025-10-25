package ingredients

import (
	"context"
	"fmt"
	"pyra/pkg/db"
	"pyra/pkg/nutrition"
	"strings"
)

type Repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) nutrition.IngredientRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) BeginTx(ctx context.Context) (db.DBTX, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *Repository) WithTx(tx db.DBTX) nutrition.IngredientRepository {
	return &Repository{
		db: tx,
	}
}

func (r *Repository) GetIngredientables(
	ctx context.Context,
	infos []nutrition.IngredientInfo,
) ([]nutrition.Ingredientable, error) {
	if len(infos) == 0 {
		return nil, nil
	}

	var productPart []nutrition.Ingredientable
	var dishPart []nutrition.Ingredientable

	var productPlaceholders []string
	var dishPlaceholders []string

	values := make([]any, 0, len(infos) * 2) // len * (UID + version)

	for i, info := range infos {
		values = append(values, info.UID, info.Version)

		placeholder := fmt.Sprintf("($%d, $%d)", i * 2 + 1, i * 2 + 2)
		ingredientable := nutrition.Ingredientable{
			Info: info,
		}


		switch info.Type {
		case nutrition.IngredientProduct:
			productPart = append(productPart, ingredientable)
			productPlaceholders = append(productPlaceholders, placeholder)
		case nutrition.IngredientDish:
			dishPart = append(dishPart, ingredientable)
			dishPlaceholders = append(dishPlaceholders, placeholder)
		default:
			panic(fmt.Errorf("unhandled ingredientable type %d", info.Type))
		}
	}

	const productsQuery = `
SELECT id, uid, version, calories, proteins, fats, carbs, %d AS ingredientable_type
FROM products
WHERE (uid, version) IN (
	%s
)`
	const dishesQuery = `
SELECT id, uid, version, calories, proteins, fats, carbs, %d AS ingredientable_type
FROM dishes
WHERE (uid, version) IN (
	%s
);`
	const union = `
UNION
`

	var query strings.Builder

	if len(productPart) > 0 {
		fmt.Fprintf(&query, productsQuery,
			nutrition.IngredientProduct, strings.Join(productPlaceholders, ",\n\t"))
	}
	
	if len(dishPart) > 0 {
		if query.Len() > 0 {
			query.WriteString(union)
		}

		fmt.Fprintf(&query, dishesQuery,
			nutrition.IngredientDish, strings.Join(dishPlaceholders, ",\n\t"))
	}

	rows, err := r.db.QueryContext(ctx, query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64 	
		var uid string
		var version int32
		var t nutrition.IngredientableType
		var m nutrition.Macro

		err = rows.Scan(&id, &uid, &version, &m.Calories, &m.Proteins, &m.Fats, &m.Carbs, &t)
		if err != nil {
			return nil, err
		}

		switch t {
		case nutrition.IngredientProduct:
			for i := range productPart {
				item := &productPart[i]
				if item.Info.UID == uid && item.Info.Version == version {
					item.ID = id
					item.Macro = m
				}
			}
		case nutrition.IngredientDish:
			for i := range dishPart {
				item := &dishPart[i]
				if item.Info.UID == uid && item.Info.Version == version {
					item.ID = id
					item.Macro = m
				}
			}
		default:
			panic("WTF?")
		}
	}

	ingredientables := make([]nutrition.Ingredientable, 0, len(productPart) + len(dishPart))
	ingredientables = append(ingredientables, productPart...)
	ingredientables = append(ingredientables, dishPart...)
	
	return ingredientables, nil
}

func (r *Repository) CreateIngredients(ctx context.Context, ingredients []nutrition.Ingredient) error {
	if len(ingredients) == 0 {
		return nil
	}

	placeholders := make([]string, 0, len(ingredients))
	values := make([]any, 0, len(ingredients)*6)

	for idx, ing := range ingredients {
		phIdx := idx * 6

		placeholders = append(placeholders, fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d)",
			phIdx + 1, phIdx + 2, phIdx + 3, phIdx + 4, phIdx + 5, phIdx + 6,
		))

		values = append(values,
			ing.DishID, ing.IngredientableID, ing.IngredientableType,
			ing.Amount, ing.Unit, ing.Idx)
	}

	query := `
INSERT INTO ingredients (
	dish_id, ingredientable_id, ingredientable_type, amount, unit, idx
) VALUES %s`

	query = fmt.Sprintf(query, strings.Join(placeholders, ",\n\t"))

	result, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != int64(len(ingredients)) {
		return fmt.Errorf("expected to insert %d ingredients, got %d", len(ingredients), affected)
	}

	return nil
}
