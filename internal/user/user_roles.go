package user

type Role string

const (
	Admin      Role = "Admin"
	Guest      Role = "Guest"
	NormalUser Role = "User"
)
