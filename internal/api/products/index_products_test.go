package products

import (
	"net/http"
	"testing"

	"pyra/test"

	"github.com/stretchr/testify/assert"
)

func Test_IndexProductsHandler(t *testing.T) {
	test.SetCWDToProjectRoot(t)

	db := test.DB(t)
	api := NewTestProductsAPI(db)
	h := test.NewMux(http.MethodGet, ProductsPATH, api.Index(), t.Output())
	
	t.Run("success", func(t *testing.T) {
		t.Cleanup(db.Truncate)

		res := h.Handle(http.MethodGet, ProductsPATH, nil)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
