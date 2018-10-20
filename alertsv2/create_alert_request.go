package alertsv2

import "net/url"

type CreateAlertRequest struct {
	Message     string            `json:"message,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Description string            `json:"description,omitempty"`
	Responders  []Recipient       `json:"responders,omitempty"`
	VisibleTo   []Recipient       `json:"visibleTo,omitempty"`
	Actions     []string          `json:"actions,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
	Entity      string            `json:"entity,omitempty"`
	Source      string            `json:"source,omitempty"`
	Priority    Priority          `json:"priority,omitempty"`
	User        string            `json:"user,omitempty"`
	Note        string            `json:"note,omitempty"`
	ApiKey      string            `json:"-"`
}

func (r *CreateAlertRequest) GenerateUrl() (string, url.Values, error) {
	return "/v2/alerts", nil, nil
}

func (r *CreateAlertRequest) GetApiKey() string {
	return r.ApiKey
}

func (r *CreateAlertRequest) Init() {
	if r.Responders != nil {
		var convertedResponder []Recipient
		for _, r := range r.Responders {
			switch r.(type) {
			case *Team:
				{
					team := r.(*Team)
					recipient := &RecipientDTO{
						Id:   team.ID,
						Name: team.Name,
						Type: "team",
					}
					convertedResponder = append(convertedResponder, recipient)
				}
			case *User:
				{
					user := r.(*User)
					recipient := &RecipientDTO{
						Id:       user.ID,
						Username: user.Username,
						Type:     "user",
					}
					convertedResponder = append(convertedResponder, recipient)
				}
			case *Escalation:
				{
					escalation := r.(*Escalation)
					recipient := &RecipientDTO{
						Id:   escalation.ID,
						Name: escalation.Name,
						Type: "escalation",
					}
					convertedResponder = append(convertedResponder, recipient)
				}
			case *Schedule:
				{
					schedule := r.(*Schedule)
					recipient := &RecipientDTO{
						Id:   schedule.ID,
						Name: schedule.Name,
						Type: "schedule",
					}
					convertedResponder = append(convertedResponder, recipient)
				}

			}
		}
		r.Responders = convertedResponder
	}

	if r.VisibleTo != nil {
		var convertedVisibleTo []Recipient
		for _, r := range r.VisibleTo {
			switch r.(type) {
			case *Team:
				{
					team := r.(*Team)
					recipient := &RecipientDTO{
						Id:   team.ID,
						Name: team.Name,
						Type: "team",
					}
					convertedVisibleTo = append(convertedVisibleTo, recipient)
				}
			case *User:
				{
					user := r.(*User)
					recipient := &RecipientDTO{
						Id:       user.ID,
						Username: user.Username,
						Type:     "user",
					}
					convertedVisibleTo = append(convertedVisibleTo, recipient)
				}
			}
		}
		r.VisibleTo = convertedVisibleTo
	}

}
