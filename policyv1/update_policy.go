package policyv1

type UpdatePolicyRequest struct {
	*Identifier
	ApiKey            string
	Name              string `json:"name"`
	Type              string `json:"type"`
	PolicyDescription string `json:"policyDescription"`
	Enabled           bool   `json:"enabled"`
	Filter            Filter `json:"filter"`
}
