package category

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	"gorm.io/gorm"
)

// ResponseToCategory takes swagger model and serialize to gorm Category model.
func ResponseToCategory(c *api.Category) *Category {
	return &Category{
		Model:     gorm.Model{},
		Name:      c.Name,
		IsDeleted: c.IsDeleted,
		IsActive:  c.IsActive,
	}
}
