package policyv1

type UpdatePolicyRequest struct {
	*Identifier
	ApiKey
	Name              string            `json:"name,omitempty"`
	Type              PolicyType        `json:"type,omitempty"`
	PolicyDescription string            `json:"policyDescription,omitempty"`
	Enabled           bool              `json:"enabled,omitempty"`
	Filter            Filter            `json:"filter,omitempty"`
	TimeRestrictions  *TimeRestrictions `json:"timeRestrictions,omitempty"`
}

type UpdateAutoClosePolicyRequest struct {
	UpdatePolicyRequest
	Duration Duration `json:"duration,omitempty"`
}

type UpdateAutoRestartPolicyRequest struct {
	UpdatePolicyRequest
	Duration       Duration `json:"duration,omitempty"`
	MaxRepeatCount int      `json:"maxRepeatCount,omitempty"`
}

type UpdateNotificationSuppressPolicyRequest struct {
	UpdatePolicyRequest
}

type UpdateNotificationDeduplicationPolicyRequest struct {
	UpdatePolicyRequest
	DeduplicationActionType DeduplicationActionType `json:"deduplicationActionType,omitempty"`
	Count                   int                     `json:"count,omitempty"`
	Duration                Duration                `json:"duration,omitempty"`
}

type UpdateNotificationDelayPolicyRequest struct {
	UpdatePolicyRequest
	DelayOption DelayOption `json:"delayOption,omitempty"`
	UntilMinute int         `json:"untilMinute,omitempty"`
	UntilHour   int         `json:"untilHour,omitempty"`
	Duration    Duration    `json:"duration,omitempty"`
}

type UpdateModifyPolicyRequest struct {
	UpdatePolicyRequest
	Message                    string      `json:"message,omitempty"`
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

// UpdatePolicyResponse is a response of updating policy
type UpdatePolicyResponse struct {
	ResponseMeta
	Status string `json:"result"`
}
