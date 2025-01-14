package dishes

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/dishes"
	"pyra/pkg/foodproducts"
)

type API struct {
	*base.API
}

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
		Handler: baseHandler,
		svc:     dishes.NewRepository(api.DB),
	}
}

func (api *API) Show() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/dishes/show.html")
	if err != nil {
		panic(err)
	}

	return &ShowDishHandler{
		Handler:      baseHandler,
		svc:          dishes.NewRepository(api.DB),
		productsRepo: foodproducts.NewRepository(api.DB),
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
