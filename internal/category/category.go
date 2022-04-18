package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name      string
	IsDeleted bool
	IsActive  bool
}

func (Category) TableName() string {
	//default table name
	return "categories"
}
