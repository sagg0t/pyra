package foodproduct

import (
	"log/slog"
	"net/http"

	view "github.com/olehvolynets/pyra/view/food_product"
)

func listFoodProducts(w http.ResponseWriter, r *http.Request) {
	// dbConf := db.NewConfig("postgres")
	// // dbConf.Attrs.Add("sslmode", fetchEnv("DB_SSLMODE", "disable"))
	// dbConn, err := pgx.Connect(r.Context(), dbConf.String())
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	//
	// rows, err := dbConn.Query(r.Context(), "SELECT * FROM food_products;")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	//
	// foodProducts := make([]foodproduct.FoodProduct, 0)
	// for rows.Next() {
	// 	fp := foodproduct.FoodProduct{}
	//
	// 	err := rows.Scan(&fp.ID, &fp.Name, &fp.Calories, &fp.Proteins, &fp.Fats, &fp.Carbs, &fp.CreatedAt, &fp.UpdatedAt)
	// 	if err != nil {
	// 		log.Printf("failed to scan a row - %v", err)
	// 	}
	//
	// 	foodProducts = append(foodProducts, fp)
	// }

	// json.NewEncoder(w).Encode(foodProducts)

	component := view.ProductList()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Warn(err.Error())
	}
}
