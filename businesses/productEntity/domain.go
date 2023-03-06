package productEntity

import (
	"context"
	"time"
)

type Domain struct {
	ID          uint
	Name        string
	Price       float64
	Description string
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ParamGetProducts struct {
	Page      int
	Offset    int
	Limit     int
	Price     int
	TypePrice string
	Type      string
	Sort      string
}

type Service interface {
	CreateNewProduct(ctx context.Context, data *Domain) (*Domain, error)
	GetProducts(ctx context.Context, params ParamGetProducts) (*[]Domain, int, int64, error)
	GetProductById(ctx context.Context, id uint) (*Domain, error)
	UpdateProductById(ctx context.Context, data *Domain, id uint) (*Domain, error)
	DeleteProductById(ctx context.Context, id uint) (string, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewProduct(ctx context.Context, data *Domain) (*Domain, error)
	GetProducts(ctx context.Context, params ParamGetProducts) (*[]Domain, int64, error)
	GetProductById(ctx context.Context, id uint) (*Domain, error)
	UpdateProductById(ctx context.Context, id uint, data *Domain) (*Domain, error)
	DeleteProductById(ctx context.Context, id uint) (string, error)
}
