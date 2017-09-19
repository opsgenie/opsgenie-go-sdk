package schedulev2

type Schedule struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Timezone string `json:"timezone"`
	Enabled bool `json:"enabled"`
	OwnerTeam OwnerTeam `json:"ownerTeam"`
	Rotations []Rotation `json:"rotations"`
}
