package foodproducts

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/foodproducts"
)

type API struct {
	*base.API
	repository *foodproducts.FoodProductsRepository
}

func NewAPI(api *base.API) *API {
	svc := foodproducts.NewRepository(api.DB)

	return &API{
		API:        api,
		repository: svc,
	}
}

func (api *API) Index() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/index.html")
	if err != nil {
		panic(err)
	}

	return &FoodProductsHandler{
		Handler: baseHandler,
		svc:     api.repository,
	}
}

func (api *API) Show() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/show.html")
	if err != nil {
		panic(err)
	}

	return &FoodProductHanler{
		Handler: baseHandler,
		svc:     api.repository,
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
		Handler: baseHandler,
		svc:     api.repository,
	}
}

func (api *API) Create() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/new.html")
	if err != nil {
		panic(err)
	}

	return &CreateFoodProductHandler{
		Handler: baseHandler,
		svc:     api.repository,
	}
}

func (api *API) Update() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/foodproducts/edit.html")
	if err != nil {
		panic(err)
	}

	return &UpdateFoodProductHandler{
		Handler: baseHandler,
		svc:     api.repository,
	}
}

func (api *API) Delete() http.Handler {
	return &DeleteFoodProductHandler{
		Handler: api.NewHandler(),
		svc:     api.repository,
	}
}

func (api *API) Search() http.Handler {
	return &SearchFoodProductsHandler{
		Handler: api.NewHandler(),
		svc:     api.repository,
	}
}
