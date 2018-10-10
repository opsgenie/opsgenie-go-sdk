package schedulev2

import (
	"fmt"
	"net/url"
	"time"
)

// GetScheduleTimelineRequest is a struct of request request the timeline of a schedule.
type GetScheduleTimelineRequest struct {
	*Identifier
	ApiKey       string
	Expand       string
	Interval     int32
	IntervalUnit string
	Date         string
}

// GetApiKey returns api key.
func (r *GetScheduleTimelineRequest) GetApiKey() string {
	return r.ApiKey
}

// GenerateUrl generates url to API endpoint.
func (r *GetScheduleTimelineRequest) GenerateUrl() (string, url.Values, error) {
	baseUrl, params, err := r.Identifier.GenerateUrl()
	if err != nil {
		return "", nil, err
	}

	baseUrl += "/timeline"

	if r.Expand != "" {
		params.Add("expand", r.Expand)
	}

	if r.Interval != 0 {
		params.Add("interval", fmt.Sprintf("%d", r.Interval))
	}

	if r.IntervalUnit != "" {
		params.Add("intervalUnit", r.IntervalUnit)
	}

	if r.Date != "" {
		params.Add("date", r.Date)
	}

	return baseUrl, params, nil
}

// GetScheduleTimelineResponse is a response of retrieving a timeline
type GetScheduleTimelineResponse struct {
	ResponseMeta

	Timeline Timeline `json:"data"`
}

// Timeline is the timeline for a specific range of dates
type Timeline struct {
	Parent        Parent        `json:"_parent,omitempty"`
	StartDate     time.Time     `json:"startDate"`
	EndDate       time.Time     `json:"endDate"`
	FinalTimeline FinalTimeline `json:"finalTimeline"`
}

// FinalTimeline contains the effective on-call rotations
type FinalTimeline struct {
	Rotations []TimelineRotation `json:"rotations"`
}

// TimelineRotation contain the timeline for a single rotation
type TimelineRotation struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Order   float64  `json:"order"`
	Periods []Period `json:"periods"`
}

// Period is a single on-call rotation
type Period struct {
	StartDate           time.Time   `json:"startDate"`
	EndDate             time.Time   `json:"endDate"`
	Type                string      `json:"type"`
	Recipient           Recipient   `json:"recipient"`
	FlattenedRecipients []Recipient `json:"flattenedRecipients"`
}

// Recipient represents a user scheduled to be on call
type Recipient struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
