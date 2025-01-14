package dishes

import (
	"fmt"
	"net/http"

	"pyra/internal/api/base"
)

type CreateDishHandler struct {
	*base.Handler
}

func (h *CreateDishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.Session(r)

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(fmt.Sprintf("failed to parse form: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(r.Form)
}

type DishForm struct {
	Name     string
	Products []struct {
		ID     string
		Amount float64
	}

	Errors map[string]string
}

func formFromParams(fetch func(key string) string) (DishForm, error) {
	form := DishForm{
		Name: fetch("name"),
	}

	return form, nil
}
