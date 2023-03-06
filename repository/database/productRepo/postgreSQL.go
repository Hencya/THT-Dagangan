package productRepo

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"THT-dagangan/businesses/productEntity"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) productEntity.Repository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateNewProduct(ctx context.Context, data *productEntity.Domain) (*productEntity.Domain, error) {
	domain := productEntity.Domain{}
	rec := Product{}
	copier.Copy(&rec, &data)
	err := r.db.Create(&rec).Error
	if err != nil {
		return nil, err
	}
	copier.Copy(&domain, &rec)
	return &domain, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context, params productEntity.ParamGetProducts) (*[]productEntity.Domain, int64, error) {
	var totalData int64
	var err error
	domain := []productEntity.Domain{}
	rec := []Product{}

	r.db.Find(&rec).Count(&totalData)
	if params.Type != "" && params.TypePrice != "" {
		switch params.TypePrice {
		case "max":
			err = r.db.Where("type = ? AND price <= ?", params.Type, params.Price).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		case "min":
			err = r.db.Where("type = ? AND price >= ?", params.Type, params.Price).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		}
	} else if params.Type == "" && params.TypePrice != "" {
		switch params.TypePrice {
		case "max":
			err = r.db.Where("price <= ?", params.Price).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		case "min":
			err = r.db.Where("price >= ?", params.Price).
				Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
				Find(&rec).Error
		}
	} else if params.Type != "" && params.TypePrice == "" {
		err = r.db.Where("type = ?", params.Type).Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
			Find(&rec).Error
	} else {
		err = r.db.Order(params.Sort).Limit(params.Limit).Offset(params.Offset).
			Find(&rec).Error
	}

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query Product: %w", err)
	}

	copier.Copy(&domain, &rec)

	return &domain, totalData, nil
}

func (r *ProductRepository) UpdateProductById(ctx context.Context, id uint, data *productEntity.Domain) (*productEntity.Domain, error) {
	domain := productEntity.Domain{}
	rec := Product{}
	recData := Product{}

	copier.Copy(&recData, &data)

	if err := r.db.First(&rec, "id = ?", id).Updates(&recData).Error; err != nil {
		fmt.Println(err)
		return &productEntity.Domain{}, err
	}

	copier.Copy(&domain, &rec)
	return &domain, nil
}

func (r *ProductRepository) GetProductById(ctx context.Context, id uint) (*productEntity.Domain, error) {
	domain := productEntity.Domain{}
	rec := Product{}

	if err := r.db.Where("id = ?", id).First(&rec).Error; err != nil {
		return &productEntity.Domain{}, err
	}
	copier.Copy(&domain, &rec)
	return &domain, nil
}

func (r *ProductRepository) DeleteProductById(ctx context.Context, id uint) (string, error) {
	rec := Product{}

	if err := r.db.Where("id = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Product was Deleted", nil
}
