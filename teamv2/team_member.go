package teamv2

// TeamMember contains team members data.
type TeamMember struct {
	User UserIdentifier `json:"user"`
	Role string         `json:"role,omitempty"`
}
