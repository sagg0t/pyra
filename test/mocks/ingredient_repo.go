package mocks

import (
	"context"

	"pyra/pkg/db"
	"pyra/pkg/nutrition"
)

type IngredientRepo struct {
	BeginTxFn            func(context.Context) (db.DBTX, error)
	WithTxFn             func(db.DBTX) nutrition.IngredientRepository
	GetIngredientablesFn func(context.Context, []nutrition.IngredientInfo) ([]nutrition.Ingredientable, error)
	CreateIngredientsFn  func(*IngredientRepo, context.Context, []nutrition.Ingredient) error
}

var _ nutrition.IngredientRepository = &IngredientRepo{}

func (mock *IngredientRepo) BeginTx(_ context.Context) (db.DBTX, error) {
	return nil, nil
}

func (mock *IngredientRepo) WithTx(_ db.DBTX) nutrition.IngredientRepository {
	return mock
}

func (mock *IngredientRepo) GetIngredientables(
	ctx context.Context,
	infos []nutrition.IngredientInfo,
) ([]nutrition.Ingredientable, error) {
	if mock.GetIngredientablesFn == nil {
		notImplemented()
	}

	return mock.GetIngredientablesFn(ctx, infos)
}

func (mock *IngredientRepo) CreateIngredients(
	ctx context.Context,
	ingredients []nutrition.Ingredient,
) error {
	if mock.CreateIngredientsFn == nil {
		notImplemented()
	}

	return mock.CreateIngredients(ctx, ingredients)
}
