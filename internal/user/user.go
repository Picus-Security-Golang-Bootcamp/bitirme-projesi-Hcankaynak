package user

import "gorm.io/gorm"

// User basic user model
type User struct {
	gorm.Model
	Id       int
	Email    string
	Password string
	Role     Role
}
