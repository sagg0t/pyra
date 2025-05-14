package nutrition

import (
	"context"
	"errors"
)

var ErrAlreadyExists = errors.New("Product already exists")

type ProductService struct {
	repo     ProductRepository
	dishRepo DishRepository
}

func NewProductService(repo ProductRepository) (*ProductService, error) {
	return &ProductService{
		repo: repo,
	}, nil
}

func (s *ProductService) List(ctx context.Context) ([]Product, error) {
	return s.repo.Index(ctx)
}

func (s *ProductService) FindByID(ctx context.Context, id ProductID) (Product, error) {
	return s.repo.FindByID(ctx, id)
}

type CreateProductInfo struct {
	UID  ProductUID
	Name ProductName
	Macro
}

func (s *ProductService) Create(
	ctx context.Context,
	info CreateProductInfo,
) (Product, error) {
	isTaken, err := s.repo.IsNameTaken(ctx, info.Name)
	if err != nil {
		return Product{}, err
	} else if isTaken {
		return Product{}, ErrAlreadyExists
	}

	product, err := s.repo.Create(
		ctx,
		info.UID,
		info.Name,
		info.Macro,
	)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

type UpdateProductInfo struct {
	ID   ProductID
	Name ProductName
	Macro
}

func (s *ProductService) Update(
	ctx context.Context,
	info UpdateProductInfo,
) (Product, error) {
	product, err := s.repo.FindByID(ctx, info.ID)
	if err != nil {
		return Product{}, err
	}

	// TODO: add TX here
	dishes, err := s.dishRepo.FindAllByProductID(ctx, info.ID)
	if err != nil {
		return Product{}, err
	}

	if len(dishes) == 0 {
		product, err := s.repo.CreateVersion(
			ctx,
			product.UID,
			info.Name,
			info.Macro,
		)
		if err != nil {
			return Product{}, err
		}

		return product, nil
	} else {
		err := s.repo.Update(ctx, info.ID, info.Name, info.Macro)
		if err != nil {
			return Product{}, err
		}

		product.Name = info.Name
		product.Macro = info.Macro

		return product, nil
	}
}
