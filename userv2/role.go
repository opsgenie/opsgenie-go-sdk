package userv2

const (
	// OwnerRole is the text value of standard role "owner"
	OwnerRole = "owner"
	// AdminRole is the text value of standard role "admin"
	AdminRole = "admin"
	// UserRole is the text value of standard role "user"
	UserRole = "user"
)

// Role contains data of role.
type Role struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
