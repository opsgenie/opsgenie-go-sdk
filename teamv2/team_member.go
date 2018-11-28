package teamv2

// TeamMember contains team members data.
type TeamMember struct {
	User string `json:"user"`
	Role string `json:"role,omitempty"`
}
