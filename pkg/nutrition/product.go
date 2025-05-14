package nutrition

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Product struct {
	ID      ProductID
	UID     ProductUID
	Version ProductVersion

	Name ProductName

	Macro

	CreatedAt time.Time
	UpdatedAt time.Time
}

type (
	ProductID      uint64
	ProductUID     string
	ProductVersion int32
	ProductName    string
)

type ProductErrors struct {
	ID      error
	UID     error
	Version error

	Name error

	MacroErrors
}

func (e *ProductErrors) HasErrors() bool {
	return e.ID != nil ||
		e.UID != nil ||
		e.Version != nil ||
		e.Name != nil ||
		e.MacroErrors.HasErrors()
}

const productErrFmt = `ID: %w
UID: %w
Version: %w
Name: %w
%s`

func (e *ProductErrors) Error() string {
	return fmt.Errorf(productErrFmt,
		e.ID, e.UID, e.Version, e.Name,
		e.MacroErrors.Error(),
	).Error()
}

var ErrNameEmpty = errors.New("name can't be empty")

func NewProductName(n string) (ProductName, error) {
	n = strings.TrimSpace(n)
	if len(n) == 0 {
		return ProductName(""), ErrNameEmpty
	}

	return ProductName(n), nil
}
