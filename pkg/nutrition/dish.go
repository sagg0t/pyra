package nutrition

import (
	"fmt"
	"strings"
	"time"
)

type Dish struct {
	ID      DishID
	UID     DishUID
	Version DishVersion

	Name DishName

	Macro

	CreatedAt time.Time
	UpdatedAt time.Time
}

type DishErrors struct {
	ID      error
	UID     error
	Version error

	Name error

	MacroErrors
}

func (e *DishErrors) HasErrors() bool {
	idErr := e.ID != nil
	uidErr := e.UID != nil
	versionErr := e.Version != nil
	nameErr := e.Version != nil

	return idErr || uidErr || versionErr || nameErr || e.MacroErrors.HasErrors()
}

func (e *DishErrors) Error() string {
	return fmt.Errorf(
		`ID: %w
		UID: %w
		Version: %w
		Name: %w
		%s`,
		e.ID, e.UID, e.Version, e.Name,
		e.MacroErrors.Error(),
	).Error()
}

type (
	DishID      uint64
	DishUID     string
	DishVersion int32
	DishName    string
)

func NewDishUID(s string) (DishUID, error) {
	return DishUID(s), nil
}

func NewDishName(n string) (DishName, error) {
	n = strings.TrimSpace(n)
	if len(n) == 0 {
		return DishName(""), ErrNameEmpty
	}

	return DishName(n), nil
}
