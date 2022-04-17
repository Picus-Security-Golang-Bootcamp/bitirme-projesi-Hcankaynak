package product

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Migration() {
	zap.L().Debug("Product Repository Migration happening.")
	r.db.AutoMigrate(&Product{})
}
