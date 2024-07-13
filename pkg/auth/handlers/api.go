package handlers

import (
	"pyra/pkg/auth"
	"pyra/pkg/pyra"
)

type API struct {
	pyra.API

	authSvc *auth.AuthService
}

func NewAPI(authSvc *auth.AuthService) *API {
	return &API{
		authSvc: authSvc,
	}
}
