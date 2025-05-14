package base

import (
	"html/template"

	"pyra/internal/users"
	"pyra/pkg/auth"
	"pyra/pkg/db"
)

type API struct {
	DB       db.DBTX
	drivers  *template.Template
	UserRepo auth.UserRepository
}

func NewAPI(db db.DBTX, drivers *template.Template) *API {
	return &API{
		DB:       db,
		drivers:  drivers,
		UserRepo: users.NewRepository(db),
	}
}

func (api *API) NewHandler() *Handler {
	baseTemplate, err := api.drivers.Clone()
	if err != nil {
		panic(err)
	}

	return &Handler{
		template: baseTemplate,
		userRepo: api.UserRepo,
	}
}
