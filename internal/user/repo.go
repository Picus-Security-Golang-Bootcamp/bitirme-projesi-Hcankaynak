package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *api.User) (*User, error) {
	zap.L().Debug("user.repo.create", zap.Reflect("userBody", u))

	user := ResponseToUser(u)
	if err := r.db.Create(&user).Error; err != nil {
		zap.L().Error("user.repo.Create failed toΩ create author", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Login(u *api.User) (*User, error) {
	zap.L().Debug("user.repo.login", zap.Reflect("userBody", u))

	var user User
	if err := r.db.Where(&User{Email: u.Email, Password: u.Password}).First(&user).Error; err != nil {
		zap.L().Error("user.repo.login failed", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) getUserByEmail(email string) (*User, error) {
	var user User
	zap.L().Debug("user.repo.get", zap.Reflect("email", email))
	if err := r.db.Where(&User{Email: email}).First(&user).Error; err != nil {
		zap.L().Debug("user.repo.Where no user found")
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindUserById(Id int) (*User, error) {
	var user User
	if err := r.db.First(&user, Id).Error; err != nil {
		zap.L().Debug("user.repo.Where no user found")
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Migration() {
	zap.L().Debug("User Repository Migration happening.")

	r.db.AutoMigrate(&User{})
}
