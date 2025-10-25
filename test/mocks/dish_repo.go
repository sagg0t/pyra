package mocks

import (
	"context"
	"pyra/pkg/db"
	"pyra/pkg/nutrition"
)

type DishRepo struct {
	BeginTxFn func(_ context.Context) (db.DBTX, error)
	WithTxFn func(_ db.DBTX) nutrition.DishRepository
	IndexFn func(context.Context) ([]nutrition.Dish, error)
	FindByIDFn func(context.Context, nutrition.DishID) (nutrition.Dish, error)
	VersionsFn func(context.Context, nutrition.DishUID) ([]nutrition.Dish, error)
	FindAllByProductIDFn func(context.Context, nutrition.ProductID) ([]nutrition.Dish, error)
	FindAllByRefsFn func(context.Context, []nutrition.DishRef) ([]nutrition.Dish, error)
	IsNameTakenFn func(context.Context, nutrition.DishName, nutrition.DishUID) (bool, error)
	CreateFn func(context.Context, *nutrition.Dish) error
}

var _ nutrition.DishRepository = &DishRepo{}

func (mock *DishRepo) BeginTx(_ context.Context) (db.DBTX, error) {
	return nil, nil
}

func (mock *DishRepo) WithTx(_ db.DBTX) nutrition.DishRepository {
	return mock
}

func (mock *DishRepo) Index(ctx context.Context) ([]nutrition.Dish, error) {
	if mock.IndexFn == nil {
		notImplemented()
	}

	return mock.IndexFn(ctx)
}

func (mock *DishRepo) FindByID(ctx context.Context, id nutrition.DishID) (nutrition.Dish, error) {
	if mock.FindByIDFn == nil {
		notImplemented()
	}

	return mock.FindByIDFn(ctx, id)
}

func (mock *DishRepo) Versions(ctx context.Context, uid nutrition.DishUID) ([]nutrition.Dish, error) {
	if mock.VersionsFn == nil {
		notImplemented()
	}

	return mock.VersionsFn(ctx, uid)
}

func (mock *DishRepo) FindAllByProductID(ctx context.Context, prodID nutrition.ProductID) ([]nutrition.Dish, error) {
	if mock.FindAllByProductIDFn == nil {
		notImplemented()
	}

	return mock.FindAllByProductIDFn(ctx, prodID)
}

func (mock *DishRepo) FindAllByRefs(ctx context.Context, refs []nutrition.DishRef) ([]nutrition.Dish, error) {
	if mock.FindAllByRefsFn == nil {
		notImplemented()
	}

	return mock.FindAllByRefsFn(ctx, refs)
}

func (mock *DishRepo) IsNameTaken(ctx context.Context, name nutrition.DishName, uid nutrition.DishUID) (bool, error) {
	if mock.IsNameTakenFn == nil {
		notImplemented()
	}

	return mock.IsNameTakenFn(ctx, name, uid)
}

func (mock *DishRepo) Create(ctx context.Context, dish *nutrition.Dish) error {
	if mock.CreateFn == nil {
		notImplemented()
	}

	return mock.CreateFn(ctx, dish)
}
