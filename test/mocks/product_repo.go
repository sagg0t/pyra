package mocks

import (
	"context"

	"pyra/pkg/db"
	"pyra/pkg/nutrition"
)

type ProductRepo struct {
	BeginTxFn      func(context.Context) (db.DBTX, error)
	WithTxFn       func(db.DBTX) nutrition.ProductRepository
	IndexFn        func(context.Context) ([]nutrition.Product, error)
	FindAllByIDsFn func(context.Context, []nutrition.ProductID) ([]nutrition.Product, error)
	ForDishFn      func(context.Context, nutrition.DishID) ([]nutrition.Product, error)

	FindByIDFn  func(context.Context, nutrition.ProductID) (nutrition.Product, error)
	FindByRefFn func(context.Context, nutrition.ProductUID, nutrition.ProductVersion) (nutrition.Product, error)

	VersionsFn func(context.Context, nutrition.ProductUID) ([]nutrition.Product, error)

	CreateFn func(context.Context, nutrition.ProductUID, nutrition.ProductName, nutrition.Macro) (nutrition.Product, error)
	CreateVersionFn func(context.Context, nutrition.ProductUID, nutrition.ProductName, nutrition.Macro) (nutrition.Product, error)
	DeleteFn        func(context.Context, nutrition.ProductID) error
	UpdateFn        func(context.Context, nutrition.ProductID, nutrition.ProductName, nutrition.Macro) error

	IsNameTakenFn func(context.Context, nutrition.ProductName) (bool, error)
	MaxVersionFn  func(context.Context, nutrition.ProductUID) (nutrition.ProductVersion, error)

	SearchFn func(ctx context.Context, searchStr string) ([]nutrition.Product, error)
}

var _ nutrition.ProductRepository = &ProductRepo{}

func (mock *ProductRepo) BeginTx(context.Context) (db.DBTX, error) {
	return nil, nil
}

func (mock *ProductRepo) WithTx(_ db.DBTX) nutrition.ProductRepository {
	return mock
}

func (mock *ProductRepo) Index(ctx context.Context) ([]nutrition.Product, error) {
	if mock.IndexFn == nil {
		notImplemented()
	}

	return mock.IndexFn(ctx)
}

func (mock *ProductRepo) FindAllByIDs(ctx context.Context, ids []nutrition.ProductID) ([]nutrition.Product, error) {
	if mock.FindAllByIDsFn == nil {
		notImplemented()
	}

	return mock.FindAllByIDsFn(ctx, ids)
}

func (mock *ProductRepo) ForDish(ctx context.Context, id nutrition.DishID) ([]nutrition.Product, error) {
	if mock.ForDishFn == nil {
		notImplemented()
	}

	return mock.ForDishFn(ctx, id)
}

func (mock *ProductRepo) FindByID(ctx context.Context, id nutrition.ProductID) (nutrition.Product, error) {
	if mock.FindByIDFn == nil {
		notImplemented()
	}

	return mock.FindByIDFn(ctx, id)
}

func (mock *ProductRepo) FindByRef(
	ctx context.Context,
	uid nutrition.ProductUID,
	version nutrition.ProductVersion,
) (nutrition.Product, error) {
	if mock.FindByRefFn == nil {
		notImplemented()
	}

	return mock.FindByRefFn(ctx, uid, version)
}


func (mock *ProductRepo) Versions(ctx context.Context, uid nutrition.ProductUID) ([]nutrition.Product, error) {
	if mock.VersionsFn == nil {
		notImplemented()
	}

	return mock.VersionsFn(ctx, uid)
}


func (mock *ProductRepo) Create(
	ctx context.Context,
	uid nutrition.ProductUID,
	name nutrition.ProductName,
	macro nutrition.Macro,
) (nutrition.Product, error) {
	if mock.CreateFn == nil {
		notImplemented()
	}

	return mock.CreateFn(ctx, uid, name, macro)
}

func (mock *ProductRepo) CreateVersion(
	ctx context.Context,
	uid nutrition.ProductUID,
	name nutrition.ProductName,
	macro nutrition.Macro,
) (nutrition.Product, error) {
	if mock.CreateVersionFn == nil {
		notImplemented()
	}

	return mock.CreateVersionFn(ctx, uid, name, macro)
}

func (mock *ProductRepo) Delete(ctx context.Context, id nutrition.ProductID) error {
	if mock.DeleteFn == nil {
		notImplemented()
	}

	return mock.DeleteFn(ctx, id)
}

func (mock *ProductRepo) Update(ctx context.Context, id nutrition.ProductID, name nutrition.ProductName, macro nutrition.Macro) error {
	if mock.UpdateFn == nil {
		notImplemented()
	}

	return mock.UpdateFn(ctx, id, name, macro)

}

func (mock *ProductRepo) IsNameTaken(ctx context.Context, name nutrition.ProductName) (bool, error) {
	if mock.IsNameTakenFn == nil {
		notImplemented()
	}

	return mock.IsNameTakenFn(ctx, name)
}

func (mock *ProductRepo) MaxVersion(ctx context.Context, uid nutrition.ProductUID) (nutrition.ProductVersion, error) {
	if mock.MaxVersionFn == nil {
		notImplemented()
	}

	return mock.MaxVersionFn(ctx, uid)
}

func (mock *ProductRepo) Search(ctx context.Context, searchStr string) ([]nutrition.Product, error) {
	if mock.SearchFn == nil {
		notImplemented()
	}

	return mock.SearchFn(ctx, searchStr)
}

