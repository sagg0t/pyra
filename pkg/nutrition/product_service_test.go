package nutrition

import (
	"testing"
)

func TestListProducts(t *testing.T) {
	t.Run("with products in repo", func(t *testing.T) {
		t.Skip("not implemented")
	})

	t.Run("without products in repo", func(t *testing.T) {
		t.Skip("not implemented")
	})
}

func TestFindProductByID(t *testing.T) {
	t.Run("product exists", func(t *testing.T) {
		t.Skip("not implemented")
	})

	t.Run("product does NOT exist", func(t *testing.T) {
		t.Skip("not implemented")
	})
}

func TestCreateProduct(t *testing.T) {
	t.Skip("not implemented")

	t.Run("with validation error", func(t *testing.T) {
		t.Skip("not implemented")
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("when product has never been used yet", func(t *testing.T) {
		t.Skip("not implemented")

		t.Run("with validation error", func(t *testing.T) {
			t.Skip("not implemented")
		})
	})

	t.Run("when product was used in a dish", func(t *testing.T) {
		t.Skip("not implemented")

		t.Run("with validation error", func(t *testing.T) {
			t.Skip("not implemented")
		})
	})
}
