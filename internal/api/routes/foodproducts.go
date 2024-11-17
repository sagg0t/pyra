package routes

import "fmt"

func EditFoodProduct(id uint64) string {
	return fmt.Sprintf("/foodProducts/%d/edit", id)
}

func FoodProduct(id uint64) string {
	return fmt.Sprintf("/foodProducts/%d", id)
}
