package basket

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BasketRepository struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) *BasketRepository {
	return &BasketRepository{db: db}
}

func (r *BasketRepository) getAll() string {
	zap.L().Debug("book.repo.getAll")
	return ""
}
