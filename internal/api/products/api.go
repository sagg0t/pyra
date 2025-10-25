package products

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/internal/dishes"
	"pyra/internal/products"
	"pyra/pkg/nutrition"
)

type API struct {
	*base.API
	ProductRepo nutrition.ProductRepository
	DishRepo    nutrition.DishRepository
}

func NewAPI(api *base.API) *API {
	repo := products.NewRepository(api.DB)
	dishRepo := dishes.NewRepository(api.DB)

	return &API{
		API:         api,
		ProductRepo: repo,
		DishRepo:    dishRepo,
	}
}

func (api *API) Index() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/products/index.html")
	if err != nil {
		panic(err)
	}

	return &ProductsHandler{
		Handler:     baseHandler,
		ProductRepo: api.ProductRepo,
	}
}

func (api *API) Show() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/products/show.html")
	if err != nil {
		panic(err)
	}

	return &ProductHandler{
		Handler:     baseHandler,
		productRepo: api.ProductRepo,
		dishRepo:    api.DishRepo,
	}
}

func (api *API) New() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/products/new.html")
	if err != nil {
		panic(err)
	}

	return &NewProductHandler{
		Handler: baseHandler,
	}
}

func (api *API) Edit() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/products/edit.html")
	if err != nil {
		panic(err)
	}

	return &EditProductHandler{
		Handler:     baseHandler,
		productRepo: api.ProductRepo,
	}
}

func (api *API) Create() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/products/new.html")
	if err != nil {
		panic(err)
	}

	return &CreateProductHandler{
		Handler:     baseHandler,
		productRepo: api.ProductRepo,
	}
}

func (api *API) Update() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/products/edit.html")
	if err != nil {
		panic(err)
	}

	return &UpdateProductHandler{
		Handler:     baseHandler,
		ProductRepo: api.ProductRepo,
		DishRepo:    api.DishRepo,
	}
}

func (api *API) Delete() http.Handler {
	return &DeleteProductHandler{
		Handler:     api.NewHandler(),
		ProductRepo: api.ProductRepo,
	}
}

func (api *API) Search() http.Handler {
	return &SearchProductsHandler{
		Handler:     api.NewHandler(),
		productRepo: api.ProductRepo,
	}
}
