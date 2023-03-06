package productRepo

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Description string  `gorm:"not null"`
	Type        string  `gorm:"not null"`
}
