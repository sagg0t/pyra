package auth

import (
	authLib "github.com/olehvolynets/pyra/pkg/auth"
	"github.com/olehvolynets/pyra/pkg/log"
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
