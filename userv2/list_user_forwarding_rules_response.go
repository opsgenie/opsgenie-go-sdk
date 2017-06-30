package userv2

import "time"

// ListUserForwardingRulesResponse is a response with list of forwarding rules.
type ListUserForwardingRulesResponse struct {
	ForwardingRules []ForwardingRule `json:"data,omitempty"`
	*ResponseMeta
}

// ForwardingRule contains data of forwarding rule.
type ForwardingRule struct {
	ID        string    `json:"id,omitempty"`
	Alias     string    `json:"alias,omitempty"`
	FromUser  FromUser  `json:"fromUser,omitempty"`
	ToUser    ToUser    `json:"toUser,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
}

// FromUser contains data of user, from whom the forwarding takes place.
type FromUser struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"username,omitempty"`
}

// FromUser contains data of user, for whom the forwarding takes place.
type ToUser struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"username,omitempty"`
}
