package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/teamv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)
	cli.SetAPIKey("96a6cff7-9399-4779-9083-e3680da91922")

	teamCli, err := cli.TeamV2()
	if err != nil {
		panic(teamCli)
	}

	response, err := teamCli.Get(teamv2.GetTeamRequest{
		Identifier: &teamv2.Identifier{Name: " 911_leader_team"},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"ID: %s\nName: %s\nDescription: %s\n",
		response.Team.ID,
		response.Team.Name,
		response.Team.Description,
	)
}
