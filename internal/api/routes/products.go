package routes

import (
	"fmt"

	"pyra/pkg/nutrition"
)

func EditProduct(product nutrition.Product) string {
	return fmt.Sprintf("/products/%s/%d/edit", product.UID, product.Version)
}

func Product(product nutrition.Product) string {
	return fmt.Sprintf("/products/%s/%d", product.UID, product.Version)
}
