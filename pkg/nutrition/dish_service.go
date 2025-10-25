package nutrition

import (
	"context"
	"errors"
	"fmt"

	"pyra/pkg/db"

	"github.com/google/uuid"
)

var ErrDishNameTaken = errors.New("Dish with such name already exists")

func ListDishes(ctx context.Context, repo DishRepository) ([]Dish, error) {
	return repo.Index(ctx)
}

type CreateDishInfo struct {
	Name DishName

	Ingredients []IngredientInfo
}

func (info *CreateDishInfo) Validate() DishErrors {
	errors := DishErrors{
		Ingredients: make([]IngredientErrors, 0, len(info.Ingredients)),
	}

	if info.Name == "" {
		errors.Name = ErrBlank
	}

	for i, ingInfo := range info.Ingredients {
		errors.Ingredients[i] = ingInfo.Validate()
	}

	return errors
}

type IngredientInfo struct {
	Idx uint64

	UID     string
	Version int32

	Type IngredientableType

	Amount float32
	Unit   MeasurementUnit
}

func (info *IngredientInfo) Validate() IngredientErrors {
	errors := IngredientErrors{}

	return errors
}

func CreateDish(
	ctx context.Context,
	repo DishRepository,
	ingredientRepo IngredientRepository,
	info CreateDishInfo,
) (dish Dish, errors DishErrors, err error) {
	dish = Dish{
		UID: DishUID(uuid.New().String()),
		Version: 1,
		Name: info.Name,
	}

	for i := range info.Ingredients {
		info.Ingredients[i].Idx = uint64(i + 1)
	}

	errors = info.Validate()
	if errors.HasErrors() {
		return dish, errors, nil
	}

	tx, err := repo.BeginTx(ctx)
	if err != nil {
		return
	}
	defer db.RollbackGuard(ctx, tx, &err)

	repo = repo.WithTx(tx)
	ingredientRepo = ingredientRepo.WithTx(tx)

	// TODO: indicate issue to user on UI.
	nameIsTaken, err := repo.IsNameTaken(ctx, info.Name, dish.UID)
	if err != nil {
		return
	} else if nameIsTaken {
		return dish, errors, ErrDishNameTaken
	}

	ingredientables, err := ingredientRepo.GetIngredientables(ctx, info.Ingredients)
	if err != nil {
		return
	}

	if len(ingredientables) != len(info.Ingredients) {
		return dish, errors, fmt.Errorf("requested %d ingredients, found only %d", len(info.Ingredients), len(ingredientables))
	}

	ingredients := make([]Ingredient, 0, len(ingredientables))

	for _, ing := range ingredientables {
		dish.Macro = dish.Macro.Add(ing.Macro)
	}

	err = repo.Create(ctx, &dish)
	if err != nil {
		return
	}

	for i := range ingredients {
		ingredients[i].DishID = dish.ID
	}

	err = ingredientRepo.CreateIngredients(ctx, ingredients)
	if err != nil {
		return
	}

	return dish, errors, tx.Commit()
}
