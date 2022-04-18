package product

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	"gorm.io/gorm"
	"strconv"
)

func ResponseToProduct(response *api.Product) *Product {
	price, _ := strconv.ParseFloat(response.Price, 32)
	return &Product{
		Model:       gorm.Model{},
		Name:        response.Name,
		Category:    response.Category,
		Description: response.Description,
		Price:       price,
		Stock:       int(response.Stock),
		SKU:         response.Sku,
		SellerId:    uint(response.SellerID),
	}
}
