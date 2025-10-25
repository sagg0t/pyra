package products

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"pyra/pkg/nutrition"

	"github.com/google/uuid"
)

var ErrNotNumber = errors.New("must be a number")

func productRef(r *http.Request) (nutrition.ProductUID, nutrition.ProductVersion, error) {
	paramUID := r.PathValue("uid")
	if _, err := uuid.Parse(paramUID); err != nil {
		return nutrition.ProductUID(""), nutrition.ProductVersion(0), err
	}
	paramVersion := r.PathValue("version")

	parsedVersion, err := strconv.ParseUint(paramVersion, 10, 64)
	if err != nil || parsedVersion == 0 {
		return nutrition.ProductUID(""), nutrition.ProductVersion(0), fmt.Errorf("invalid version: %s", paramVersion)
	}

	return nutrition.ProductUID(paramUID), nutrition.ProductVersion(parsedVersion), nil
}
