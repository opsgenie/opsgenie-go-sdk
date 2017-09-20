package policyv1

// CreatePolicyRequest is a request to create new policy
type CreatePolicyRequest struct {
	*Identifier
	ApiKey            string
	Name              string            `json:"name"`
	Type              PolicyType        `json:"type"`
	PolicyDescription string            `json:"policyDescription,omitempty"`
	Enabled           bool              `json:"enabled"`
	Filter            Filter            `json:"filter,omitempty"`
	TimeRestrictions  *TimeRestrictions `json:"timeRestrictions,omitempty"`
}

// CreateAutoClosePolicyRequest is a request for creating auto close policy
type CreateAutoClosePolicyRequest struct {
	CreatePolicyRequest
	Duration Duration `json:"duration"`
}

// CreateAutoRestartPolicyRequest is a request for creating auto restart policy
type CreateAutoRestartPolicyRequest struct {
	CreatePolicyRequest
	Duration       Duration `json:"duration"`
	MaxRepeatCount int      `json:"maxRepeatCount"`
}

// CreateNotificationSuppressPolicyRequest is a request for creating notification suppress policy
type CreateNotificationSuppressPolicyRequest struct {
	CreatePolicyRequest
}

// CreateNotificationDeduplicationPolicyRequest is a request for creating notification deduplication policy
type CreateNotificationDeduplicationPolicyRequest struct {
	CreatePolicyRequest
	DeduplicationActionType DeduplicationActionType `json:"deduplicationActionType"`
	Count                   int                     `json:"count"`
	Duration                Duration                `json:"duration,omitempty"`
}

// CreateNotificationDelayPolicyRequest is a request for creating notification delay policy
type CreateNotificationDelayPolicyRequest struct {
	CreatePolicyRequest
	DelayOption DelayOption `json:"delayOption"`
	UntilMinute int         `json:"untilMinute"`
	UntilHour   int         `json:"untilHour"`
	Duration    Duration    `json:"duration,omitempty"`
}

// CreateModifyPolicyRequest is a request for creating modify policy
type CreateModifyPolicyRequest struct {
	CreatePolicyRequest
	Message                    string      `json:"message"`
	Continue                   bool        `json:"continue,omitempty"`
	Alias                      string      `json:"alias,omitempty"`
	Description                string      `json:"description,omitempty"`
	Entity                     string      `json:"entity,omitempty"`
	Source                     string      `json:"source,omitempty"`
	IgnoreOriginalAlertActions bool        `json:"ignoreOriginalAlertActions,omitempty"`
	AlertActions               string      `json:"alertActions,omitempty"`
	IgnoreOriginalDetails      bool        `json:"ignoreOriginalDetails,omitempty"`
	Details                    []string    `json:"details,omitempty"`
	IgnoreOriginalRecipients   bool        `json:"ignoreOriginalRecipients,omitempty"`
	Recipients                 []Recipient `json:"recipients,omitempty"`
	IgnoreOriginalTags         bool        `json:"ignoreOriginalTags,omitempty"`
	Tags                       []string    `json:"tags,omitempty"`
	IgnoreOriginalTeams        bool        `json:"ignoreOriginalTeams,omitempty"`
	Priority                   Priority    `json:"priority,omitempty"`
}

// GetApiKey return api key
func (r *CreatePolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// CreatePolicyResponse is a response of creating new policy
type CreatePolicyResponse struct {
	ResponseMeta
	Policy CreatePolicyResult `json:"data"`
}

// CreatePolicyResult is struct, which contains only base attributes of policy
type CreatePolicyResult struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}
