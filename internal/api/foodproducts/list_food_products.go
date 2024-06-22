package foodproducts

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"

	"github.com/olehvolynets/pyra/internal/db"
	"github.com/olehvolynets/pyra/internal/foodproducts"
	view "github.com/olehvolynets/pyra/view/foodproducts"
)

func listFoodProducts(w http.ResponseWriter, r *http.Request) {
	dbConf := db.NewConfig("postgres")
	dbConn, err := pgx.Connect(r.Context(), dbConf.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	rows, err := dbConn.Query(r.Context(), "SELECT * FROM food_products;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	foodProducts := make([]foodproducts.FoodProduct, 0)
	for rows.Next() {
		fp := foodproducts.FoodProduct{}

		err := rows.Scan(&fp.ID, &fp.Name, &fp.Calories, &fp.Proteins, &fp.Fats, &fp.Carbs, &fp.CreatedAt, &fp.UpdatedAt)
		if err != nil {
			log.Printf("failed to scan a row - %v", err)
		}

		foodProducts = append(foodProducts, fp)
	}

	component := view.ProductList(foodProducts)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Warn(err.Error())
	}
}
