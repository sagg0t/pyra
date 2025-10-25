package products

import (
	"net/http"
	"testing"

	"pyra/test"

	"github.com/stretchr/testify/assert"
)

func Test_NewProductHandler(t *testing.T) {
	test.SetCWDToProjectRoot(t)

	db := test.DB(t)
	api := NewTestProductsAPI(db)
	h := test.NewMux(http.MethodGet, NewProductPATH, api.New(), t.Output())

	t.Run("success", func(t *testing.T) {
		t.Cleanup(db.Truncate)

		res := h.Handle(http.MethodGet, NewProductURI(), nil)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
