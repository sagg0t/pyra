package routes

import (
	"fmt"

	"pyra/pkg/nutrition"
)

func EditFoodProduct(id nutrition.ProductID) string {
	return fmt.Sprintf("/foodProducts/%d/edit", uint64(id))
}

func FoodProduct(id nutrition.ProductID) string {
	return fmt.Sprintf("/foodProducts/%d", uint64(id))
}
