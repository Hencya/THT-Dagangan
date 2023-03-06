package repository

import (
	"gorm.io/gorm"

	"THT-dagangan/businesses/productEntity"
	"THT-dagangan/repository/database/productRepo"
)

func NewProductRepository(db *gorm.DB) productEntity.Repository {
	return productRepo.NewProductRepository(db)
}
