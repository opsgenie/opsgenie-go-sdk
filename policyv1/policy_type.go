package policyv1

const (
	AutoRestartNotificationsPolicyType  PolicyType = "auto-restart-notifications"
	AutoClosePolicyType                 PolicyType = "auto-close"
	NotificationSuppressPolicyType      PolicyType = "notification-suppress"
	NotificationDeduplicationPolicyType PolicyType = "notification-deduplication"
	NotificationDelayPolicyType         PolicyType = "notification-delay"
	ModifyPolicyType                    PolicyType = "modify"
)

// PolicyType is a type of the policy, which should be one of "auto-restart-notifications", "auto-close",
// "notification-suppress", "notification-deduplication", "notification-delay", "modify"
type PolicyType string
