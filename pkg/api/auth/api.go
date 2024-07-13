package auth

import (
	authLib "pyra/pkg/auth"
	"pyra/pkg/log"
)

type API struct {
	log     *log.Logger
	authSvc *authLib.AuthService
}

func NewAPI(logger *log.Logger, authSvc *authLib.AuthService) *API {
	return &API{
		log:     logger,
		authSvc: authSvc,
	}
}
