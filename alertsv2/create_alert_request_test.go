package alertsv2_test

import (
	"encoding/json"
	"testing"

	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
)

func TestCreateAlertRequest(t *testing.T) {

	responders := []alertsv2.Recipient{
		&alertsv2.Escalation{ID: "escalationId"},
		&alertsv2.Escalation{Name: "escalationName"},
		&alertsv2.Schedule{ID: "scheduleId"},
		&alertsv2.Schedule{Name: "scheduleName"},
		&alertsv2.Team{ID: "teamId"},
		&alertsv2.Team{Name: "teamName"},
		&alertsv2.User{ID: "userId"},
		&alertsv2.User{Username: "user@opsgenie.com"},
	}

	visibleTo := []alertsv2.Recipient{
		&alertsv2.Team{ID: "teamId"},
		&alertsv2.Team{Name: "teamName"},
		&alertsv2.User{ID: "userId"},
		&alertsv2.User{Username: "user@opsgenie.com"},
	}

	createAlertRequest := alertsv2.CreateAlertRequest{
		Message:     "message",
		Alias:       "alias",
		Description: "alert description",
		Responders:  responders,
		VisibleTo:   visibleTo,
		Actions:     []string{"action1", "action2"},
		Tags:        []string{"tag1", "tag2"},
		Details: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		Entity:   "entity",
		Source:   "source",
		Priority: alertsv2.P1,
		User:     "user@opsgenie.com",
		Note:     "alert note",
	}

	createAlertRequest.Init()

	type expectedFormat struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Type     string `json:"type"`
	}
	expectations := []expectedFormat{
		expectedFormat{Type: "escalation", ID: "escalationId"},
		expectedFormat{Type: "escalation", Name: "escalationName"},
		expectedFormat{Type: "schedule", ID: "scheduleId"},
		expectedFormat{Type: "schedule", Name: "scheduleName"},
		expectedFormat{Type: "team", ID: "teamId"},
		expectedFormat{Type: "team", Name: "teamName"},
		expectedFormat{Type: "user", ID: "userId"},
		expectedFormat{Type: "user", Username: "user@opsgenie.com"},
	}

	for _, responder := range createAlertRequest.Responders {
		// JSON marshal and unmarshal to ensure field names
		responderJSON, _ := json.Marshal(responder)

		var actualResponder expectedFormat
		err := json.Unmarshal(responderJSON, &actualResponder)
		if err != nil {
			t.Error(err)
			t.Fatal("Expected no error unmarshalling response, but there was one.")
		}

		var ok bool
		for _, expectedResponder := range expectations {
			if actualResponder == expectedResponder {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatal("Expected to match responder, but it didn't")
		}

	}

}
