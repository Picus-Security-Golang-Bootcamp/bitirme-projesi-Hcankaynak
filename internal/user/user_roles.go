package user

type Role string

const (
	Admin      Role = "Admin"
	Guest      Role = "Guest"
	NormalUser Role = "User"
)

func (r Role) toString() string {
	return string(r)
}
