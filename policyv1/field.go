package policyv1

const (
	MessageField         Field = "message"
	AliasField           Field = "alias"
	DescriptionField     Field = "description"
	SourceField          Field = "source"
	EntityField          Field = "entity"
	TagsField            Field = "tags"
	ActionsField         Field = "actions"
	ExtraPropertiesField Field = "extra-properties"
	RecipientsField      Field = "recipients"
	TeamsField           Field = "teams"
	PriorityField        Field = "priority"
)

type Field string
