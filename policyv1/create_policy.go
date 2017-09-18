package policyv1

// CreatePolicyRequest is a request to create new policy
type CreatePolicyRequest struct {
	*Identifier
	ApiKey            string
	Name              string            `json:"name"`
	Type              PolicyType        `json:"type"`
	PolicyDescription string            `json:"policyDescription"`
	Enabled           bool              `json:"enabled"`
	Filter            Filter            `json:"filter"`
	TimeRestrictions  *TimeRestrictions `json:"timeRestrictions,omitempty"`
}

// CreateAutoClosePolicyRequest is a request for creating auto close policy
type CreateAutoClosePolicyRequest struct {
	CreatePolicyRequest
	Duration Duration `json:"duration,omitempty"`
}

// CreateAutoRestartPolicyRequest is a request for creating auto restart policy
type CreateAutoRestartPolicyRequest struct {
	CreatePolicyRequest
	Duration       Duration `json:"duration,omitempty"`
	MaxRepeatCount int      `json:"maxRepeatCount"`
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
	Message                    string   `json:"message"`
	Continue                   bool     `json:"continue"`
	Alias                      string   `json:"alias"`
	Description                string   `json:"description"`
	Entity                     string   `json:"entity"`
	Source                     string   `json:"source"`
	IgnoreOriginalAlertActions bool     `json:"ignoreOriginalAlertActions"`
	AlertActions               string   `json:"alertActions"`
	IgnoreOriginalDetails      bool     `json:"ignoreOriginalDetails"`
	Details                    string   `json:"details"`
	IgnoreOriginalRecipients   bool     `json:"ignoreOriginalRecipients"`
	Recipients                 []string `json:"recipients"`
	IgnoreOriginalTags         bool     `json:"ignoreOriginalTags"`
	Tags                       []string `json:"tags"`
	IgnoreOriginalTeams        bool     `json:"ignoreOriginalTeams"`
	Priority                   Priority `json:"priority"`
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
