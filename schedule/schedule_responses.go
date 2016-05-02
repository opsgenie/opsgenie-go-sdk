package schedule

// Create escalation response structure
type CreateScheduleResponse struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Code int `json:"code"`
}

// Update escalation response structure
type UpdateScheduleResponse struct {
	Status string `json:"status"`
        Code int `json:"code"`
}

// Delete escalation response structure
type DeleteScheduleResponse struct {
	Status string `json:"status"`
        Code int `json:"code"`
}

// Participant
type Participant struct {
	Participant string `json:"participant,omitempty"`
	Type string `json:"type,omitempty"`
}

// RotationInfo defines the structure for each rotation definition
type RotationInfo struct {
	Id string `json:"id,omitempty"`
        StartDate string `json:"startDate,omitempty"`
        EndDate string `json:"endDate,omitempty"`
        RotationType string `json:"rotationType,omitempty"`
        Participants []Participant `json:"participants,omitempty"`
        Name string `json:"name,omitempty"`
        RotationLength int `json:"rotationLength,omitempty"`
        Restrictions []Restriction `json:"restrictions,omitempty"`
}

// Get escalation structure
type GetScheduleResponse struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Team string `json:"team,omitempty"`
	Rules []RotationInfo `json:"members,omitempty"`
}

// List escalations response structure
type ListSchedulesResponse struct {
	Schedules []GetScheduleResponse `json:"schedules,omitempty"`
}
