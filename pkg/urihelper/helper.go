package urihelper

import "fmt"

type uriKey string

const (
	product       uriKey = "product"
	products      uriKey = "product"
	newProduct    uriKey = "product"
	createProduct uriKey = "product"
	editProduct   uriKey = "product"
)

var Helper helper

type helper struct {
	formats map[uriKey]string
}

func (h *helper) ProductURI(id uint64) string {
	return h.getURI(product)
}

func (h *helper) ProductsURI() string {
	return h.getURI(products)
}

func (h *helper) NewProductURI() string {
	return h.getURI(newProduct)
}

func (h *helper) getURI(name uriKey, uriArgs ...any) string {
	format, ok := h.formats[name]
	if !ok {
		panic(fmt.Sprintf("missing URI format %s", name))
	}

	return fmt.Sprintf(format, uriArgs...)
}
