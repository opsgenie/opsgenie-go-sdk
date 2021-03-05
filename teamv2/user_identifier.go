package teamv2

// UserIdentifier identifies a user in a team.
type UserIdentifier struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"username,omitempty"`
}
