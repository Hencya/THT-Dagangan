package productEntity

import (
	"context"
	"time"

	"THT-dagangan/businesses"
)

type ProductServices struct {
	ProductRepository Repository
	ContextTimeout    time.Duration
}

func NewProductServices(repoProduct Repository, timeout time.Duration) Service {
	return &ProductServices{
		ProductRepository: repoProduct,
		ContextTimeout:    timeout,
	}
}

func (s *ProductServices) CreateNewProduct(ctx context.Context, data *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	res, err := s.ProductRepository.CreateNewProduct(ctx, data)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	return res, nil
}

func (s *ProductServices) GetProducts(ctx context.Context, params ParamGetProducts) (*[]Domain, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	if params.Page == 1 {
		params.Offset = 0
	} else {
		params.Offset = (params.Page - 1) * params.Limit
	}

	res, totalData, err := s.ProductRepository.GetProducts(ctx, params)
	if err != nil {
		return &[]Domain{}, -1, -1, businesses.ErrNotFoundProduct
	}

	return res, params.Offset, totalData, nil
}

func (s *ProductServices) UpdateProductById(ctx context.Context, data *Domain, id uint) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	dataUpdated, err := s.ProductRepository.UpdateProductById(ctx, id, data)
	if err != nil {
		return &Domain{}, businesses.ErrInternalServer
	}
	return dataUpdated, nil
}

func (s *ProductServices) DeleteProductById(ctx context.Context, id uint) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	res, err := s.ProductRepository.DeleteProductById(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundProduct
	}
	return res, nil
}

func (s *ProductServices) GetProductById(ctx context.Context, id uint) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	result, err := s.ProductRepository.GetProductById(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}
