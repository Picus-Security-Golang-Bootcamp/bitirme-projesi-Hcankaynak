package product

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	"gorm.io/gorm"
)

func ResponseToProduct(response *api.Product) *Product {
	return &Product{
		Model:       gorm.Model{},
		Name:        response.Name,
		Category:    "",
		Description: response.Description,
		Price:       0,
		Stock:       0,
		SKU:         "",
		SellerId:    0,
	}
}
