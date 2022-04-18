package basket

import (
	"errors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/product"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BasketRepository struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) *BasketRepository {
	return &BasketRepository{db: db}
}

func (r *BasketRepository) getBasketByUserId(userId int) (*[]product.Product, error) {
	zap.L().Debug("basket.repo.getByUserId", zap.Reflect("userId", userId))
	var productList []product.Product
	if err := r.db.Where(&Basket{UserID: userId}).Find(&productList); err != nil {
		return nil, errors.New("no basket found")
	}
	return &productList, nil
}

func (r *BasketRepository) createBasketItem(basket Basket) error {
	zap.L().Debug("basket.repo.createBasketItem", zap.Reflect("basketItem", basket))
	if err := r.db.Create(&basket).Error; err != nil {
		zap.L().Error("basket.repo.createBasketItem", zap.Error(err))
		return err
	}
	return nil
}

func (r *BasketRepository) Migration() {
	r.db.AutoMigrate(&Basket{})
}
