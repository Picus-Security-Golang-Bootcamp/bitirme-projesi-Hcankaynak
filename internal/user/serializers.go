package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	"gorm.io/gorm"
)

func UserToResponse(u *User) *api.User {
	return &api.User{
		Email:    u.Email,
		ID:       uint64(u.ID),
		Name:     u.Name,
		Password: u.Password,
		Role:     u.Role.toString(),
	}
}

func ResponseToUser(u *api.User) *User {
	return &User{
		Model:    gorm.Model{},
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     Role(u.Role),
	}
}
