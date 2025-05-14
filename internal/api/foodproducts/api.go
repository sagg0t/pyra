package foodproducts

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/internal/dishes"
	"pyra/internal/foodproducts"
	"pyra/pkg/nutrition"
)

type API struct {
	*base.API
	productRepo nutrition.ProductRepository
	dishRepo    nutrition.DishRepository
}

func NewAPI(api *base.API) *API {
	repo := foodproducts.NewRepository(api.DB)
	dishRepo := dishes.NewRepository(api.DB)

	return &API{
		API:         api,
		productRepo: repo,
		dishRepo:    dishRepo,
	}
}

func (api *API) Index() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/index.html")
	if err != nil {
		panic(err)
	}

	return &FoodProductsHandler{
		Handler:     baseHandler,
		productRepo: api.productRepo,
	}
}

func (api *API) Show() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/show.html")
	if err != nil {
		panic(err)
	}

	return &FoodProductHandler{
		Handler:     baseHandler,
		productRepo: api.productRepo,
		dishRepo:    api.dishRepo,
	}
}

func (api *API) New() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/new.html")
	if err != nil {
		panic(err)
	}

	return &NewFoodProductHandler{
		Handler: baseHandler,
	}
}

func (api *API) Edit() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/edit.html")
	if err != nil {
		panic(err)
	}

	return &EditFoodProductHandler{
		Handler:     baseHandler,
		productRepo: api.productRepo,
	}
}

func (api *API) Create() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/new.html")
	if err != nil {
		panic(err)
	}

	return &CreateFoodProductHandler{
		Handler:     baseHandler,
		productRepo: api.productRepo,
	}
}

func (api *API) Update() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/edit.html")
	if err != nil {
		panic(err)
	}

	return &UpdateFoodProductHandler{
		Handler:     baseHandler,
		productRepo: api.productRepo,
	}
}

func (api *API) Delete() http.Handler {
	return &DeleteFoodProductHandler{
		Handler:     api.NewHandler(),
		productRepo: api.productRepo,
	}
}

func (api *API) Search() http.Handler {
	return &SearchFoodProductsHandler{
		Handler:     api.NewHandler(),
		productRepo: api.productRepo,
	}
}
