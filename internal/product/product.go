package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Category    string
	Description string
	Price       float64
	Stock       int
	SKU         string
	SellerId    uint
}

func (Product) TableName() string {
	//default table name
	return "Products"
}
