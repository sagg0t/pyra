package nutrition

import (
	"fmt"
	"time"
)

type Product struct {
	ProductRecord

	Errors ProductErrors
}

func (p *Product) HasErrors() bool {
	return p.Errors.HasErrors()
}

func (p *Product) IsArchived() bool {
	return !p.ArchivedAt.IsZero()
}

type ProductRecord struct {
	ID      ProductID `fake:"-"`
	UID     ProductUID `fake:"{uuid}"`
	Version ProductVersion `fake:"1"`

	Name ProductName `fake:"{productname}"`

	Macro

	ArchivedAt time.Time `fake:"-"`

	CreatedAt time.Time `fake:"-"`
	UpdatedAt time.Time `fake:"-"`
}

func (p *ProductRecord) Format(f fmt.State, verb rune) {
	f.Write([]byte("aasdlfkj"))
}

func (p *ProductRecord) String() string {
	return ""
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

func (name ProductName) Validate() error {
	if len(name) == 0 {
		return ErrBlank
	}

	return nil
}
