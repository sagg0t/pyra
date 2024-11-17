package routes

import "fmt"

func DishURI(id uint64) string {
	return fmt.Sprintf("/dishes/%d", id)
}
