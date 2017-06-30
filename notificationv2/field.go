package notificationv2

const (
	// The list of fields, which are used for build filter of alerts.
	ActionsField         = "actions"
	AliasField           = "alias"
	DescriptionField     = "description"
	EntityField          = "entity"
	MessageField         = "message"
	RecipientsField      = "recipients"
	SourceField          = "source"
	TeamsField           = "teams"
	ExtraPropertiesField = "extra-properties"
)

// Field is the name of alert field, which is used to build filter of alerts.
type Field string
