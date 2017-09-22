package notificationv2

import (
	"errors"
	"net/url"
)

// Identifier defined the set of attributes for identification notification.
type Identifier struct {
	UserID       string       `json:"-"`
	RuleID       string       `json:"-"`
	StatusAction StatusAction `json:"-"`
}

// GenerateUrl generates API url using specified attributes of identifier.
func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	if request.UserID == "" {
		return "", nil, errors.New("UserID should be provided")
	}

	if len(request.StatusAction) != 0 && request.RuleID == "" {
		return "", nil, errors.New("RuleID should be provided along with disabled/enabled action")
	}

	baseUrl := "/v2/users/" + request.UserID + "/notification-rules"

	if request.RuleID != "" {
		baseUrl += "/" + request.RuleID
	}

	if request.StatusAction == EnableStatusAction {
		baseUrl += "/" + EnableStatusAction
	}

	if request.StatusAction == DisableStatusAction {
		baseUrl += "/" + DisableStatusAction
	}

	return baseUrl, nil, nil
}
