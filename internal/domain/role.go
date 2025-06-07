package domain

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
	RoleModerator Role = "moderator"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleModerator, RoleUser:
		return true
	default:
		return false
	}
}
