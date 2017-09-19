package schedulev2

const (
	UserParticipantType       ParticipantType = "user"
	EscalationParticipantType ParticipantType = "escalation"
	TeamParticipantType       ParticipantType = "team"
	NoneParticipantType       ParticipantType = "none"
)

type ParticipantType string
