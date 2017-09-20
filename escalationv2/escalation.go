package escalationv2

type Escalation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rules       []Rule `json:"rules"`
}
