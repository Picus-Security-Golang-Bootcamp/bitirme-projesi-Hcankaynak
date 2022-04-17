package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Category    string
	Description string
	Price       float32
	Stock       int
	SKU         string
	SellerId    uint
}

func (Product) TableName() string {
	//default table name
	return "products"
}
