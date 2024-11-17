package base

import (
	"html/template"

	"github.com/jackc/pgx/v5/pgxpool"

	"pyra/pkg/users"
)

type API struct {
	DB      *pgxpool.Pool
	drivers *template.Template
	UserSvc *users.UserRepository
}

func NewAPI(db *pgxpool.Pool, drivers *template.Template) *API {
	return &API{
		DB:      db,
		drivers: drivers,
		UserSvc: users.NewRepository(db),
	}
}

func (api *API) NewHandler() *Handler {
	baseTemplate, err := api.drivers.Clone()
	if err != nil {
		panic(err)
	}

	return &Handler{
		template: baseTemplate,
		userSvc:  api.UserSvc,
	}
}
