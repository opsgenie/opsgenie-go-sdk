package main

import (
	"time"
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/schedulev2"
)

func main() {
	cli := new(client.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	scheduleCli, err := cli.ScheduleV2()
	if err != nil {
		panic(err)
	}

	ownerTeam := schedulev2.OwnerTeam{
		Name: "our_team_test",
		ID: "1",
	}

	participants := []schedulev2.Participant{
		{
			Type: schedulev2.UserParticipantType,
			Username: "test@test.com",
		},
	}

	var (
		startHour = 3
		endHour = 7
		startMin = 0
		endMin = 30
	)

	timeRestriction := schedulev2.TimeRestriction{
		Type: schedulev2.TimeOfDayTimeRestrictionType,
		Restriction: schedulev2.Restriction{
			StartHour: &startHour,
			EndHour: &endHour,
			StartMin: &startMin,
			EndMin: &endMin,
		},
	}

	rotations := []schedulev2.Rotation{
		{
			Name: "rotation_of_test_escalation",
			StartDate: time.Now(),
			EndDate: time.Now().Add(time.Hour),
			Type: schedulev2.HourlyRotationType,
			Length: 1,
			Participants: participants,
			TimeRestriction: timeRestriction,
		},
	}

	response, err := scheduleCli.Create(schedulev2.CreateScheduleRequest{
		Name: "our_tested_escalation",
		Description: "it is test escalation",
		Timezone: "Europe/Kirov",
		Enabled: true,
		OwnerTeam: ownerTeam,
		Rotations: rotations,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("ID: %s\nName: %s\nEnabled: %t\n", response.Result.ID, response.Result.Name, response.Result.Enabled)
}

