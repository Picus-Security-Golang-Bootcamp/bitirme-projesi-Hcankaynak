package user

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Migration() {
	zap.L().Debug("User Repository Migration happening.")

	r.db.AutoMigrate(&User{})
}
