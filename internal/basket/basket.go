package basket

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/product"
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	UserID  int
	Product product.Product `gorm:"foreignKey:UserID"`
}

func (Basket) TableName() string {
	//default table name
	return "Baskets"
}
