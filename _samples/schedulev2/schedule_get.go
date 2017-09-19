package main

import (
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

	response, err := scheduleCli.Get(schedulev2.GetScheduleRequest{
		Identifier: &schedulev2.Identifier{ID: "1"},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(
			"ID: %s\nName: %s\nDescription: %s\nTimezone: %s\nEnabled: %t\n",
			response.Schedule.ID,
			response.Schedule.Name,
			response.Schedule.Description,
			response.Schedule.Timezone,
			response.Schedule.Enabled,
		)
		fmt.Printf("\nOwner Team: %s\n", response.Schedule.OwnerTeam.Name)
		fmt.Printf("\nRotations:\n")
		for _, rotation := range response.Schedule.Rotations {
			fmt.Printf("name: %s, type: %s\n", rotation.Name, rotation.Type)
		}
	}
}
