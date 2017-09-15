package policyv1

const (
	AutoRestartNotificationsPolicyType  PolicyType = "auto-restart-notifications"
	AutoClosePolicyType                 PolicyType = "auto-close"
	NotificationSuppressPolicyType      PolicyType = "notification-suppress"
	NotificationDeduplicationPolicyType PolicyType = "notification-deduplication"
	NotificationDelayPolicyType         PolicyType = "notification-delay"
	ModifyPolicyType                    PolicyType = "modify"
)

type PolicyType string
