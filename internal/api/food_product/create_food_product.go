package foodproduct

import (
	"fmt"
	"net/http"
)

func createFoodProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created!")
}
