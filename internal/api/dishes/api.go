package dishes

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/internal/dishes"
	"pyra/internal/products"
)

type API struct {
	*base.API
}

// NewAPI - creates API instance for dishes-related endpoints.
func NewAPI(api *base.API) *API {
	return &API{
		API: api,
	}
}

func (api *API) Index() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/dishes/index.html")
	if err != nil {
		panic(err)
	}

	return &ListDishesHandler{
		Handler:  baseHandler,
		dishRepo: dishes.NewRepository(api.DB),
	}
}

func (api *API) Show() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/dishes/show.html")
	if err != nil {
		panic(err)
	}

	return &ShowDishHandler{
		Handler:     baseHandler,
		dishRepo:    dishes.NewRepository(api.DB),
		productRepo: products.NewRepository(api.DB),
	}
}

func (api *API) New() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/dishes/new.html")
	if err != nil {
		panic(err)
	}

	return &NewDishHandler{
		Handler: baseHandler,
	}
}

func (api *API) Create() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/dishes/new.html")
	if err != nil {
		panic(err)
	}

	return &CreateDishHandler{
		Handler: baseHandler,
	}
}
