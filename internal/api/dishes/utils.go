package dishes

import (
	"net/http"
	"strconv"

	"pyra/pkg/dishes"
)

func dishID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}

type DishForm struct {
	dishes.Dish

	Errors map[string]string
}
