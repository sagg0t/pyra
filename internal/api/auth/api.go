package auth

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/auth"
	"pyra/pkg/users"
)

type API struct {
	*base.API
	svc *auth.AuthService
}

func NewAPI(api *base.API) *API {
	svc := auth.NewService(
		api.DB,
		auth.NewProviderRepository(api.DB),
		users.NewRepository(api.DB),
	)

	return &API{
		API: api,
		svc: svc,
	}
}

func (api *API) SignIn() http.Handler {
	baseHandler := api.NewHandler()
	err := baseHandler.ExpandTemplate("view/auth/sign_in.html")
	if err != nil {
		panic(err)
	}

	return &SignInHandler{
		Handler: baseHandler,
	}
}

func (api *API) SignOut() http.Handler {
	return &SignOutHandler{
		Handler: api.NewHandler(),
	}
}

func (api *API) GoogleAuthorize() http.Handler {
	return &GoogleAuthHandler{
		Handler: api.NewHandler(),
	}
}

func (api *API) GoogleCallback() http.Handler {
	return &GoogleCallbackHandler{
		Handler: api.NewHandler(),
		authSvc: api.svc,
	}
}
