package user

import "gorm.io/gorm"

// User basic user model
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     Role
}

func (User) TableName() string {
	//default table name
	return "Users"
}
