package category

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository create Category Repository and return it.
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// createCategory creating individual categories
func (r *CategoryRepository) createCategory(c *api.Category) (*Category, error) {
	zap.L().Debug("category.repo.create", zap.Reflect("categoryBody", c))

	cat := ResponseToCategory(c)
	if err := r.db.Create(&cat).Error; err != nil {
		zap.L().Error("user.repo.Create failed toΩ create author", zap.Error(err))
		return nil, err
	}

	return cat, nil
}

// GetAllCategories getting all categories if not deleted and still active.
func (r *CategoryRepository) GetAllCategories(pageIndex, pageSize int) ([]Category, int) {
	zap.L().Debug("category.repo.all")
	var categories = &[]Category{}
	var count int64
	// TODO: Find condition not working.
	r.db.Offset((pageIndex-1)*pageSize).Limit(pageSize).Find(&categories, "IsDeleted = ?").Count(&count)

	return *categories, int(count)
}

func (r *CategoryRepository) Migration() {
	r.db.AutoMigrate(&Category{})
}
