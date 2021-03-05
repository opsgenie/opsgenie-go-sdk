package teamv2

// Team contains team data.
type Team struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Members     []TeamMember `json:"members,omitempty"`
}
