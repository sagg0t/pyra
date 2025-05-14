package dishes

import (
	"net/http"
	"strconv"

	"pyra/pkg/nutrition"
)

func dishID(r *http.Request) (nutrition.DishID, error) {
	paramID := r.PathValue("id")
	uintID, err := strconv.ParseUint(paramID, 10, 64)

	return nutrition.DishID(uintID), err
}
