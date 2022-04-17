package user

// User basic user model
type User struct {
	Id       int
	Email    string
	Password string
	Roles    Role
}
