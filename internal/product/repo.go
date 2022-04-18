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
	r.db.AutoMigrate(&Product{})
}

func (r *ProductRepository) FindById(Id int) (*Product, error) {
	var product Product
	if err := r.db.First(&product, Id).Error; err != nil {
		zap.L().Debug("product.repo.Where no product found")
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) createProduct(product Product) (*Product, error) {
	zap.L().Debug("product.repo.createProduct", zap.Reflect("product", product))
	if err := r.db.Create(&product).Error; err != nil {
		zap.L().Error("product.repo.createProduct failed", zap.Error(err))
		return nil, err
	}
	return &product, nil
}
