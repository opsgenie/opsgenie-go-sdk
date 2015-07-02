package main

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

func main() {
	API_KEY := "YOUR API KEY HERE"
	TEAM_NAME := "YOUR TEAM NAME HERE"

	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// assign the owner for the alert
	addTeamReq := alerts.AddTeamAlertRequest{Id: response.AlertId, Team: TEAM_NAME}
	addTeamResponse, alertErr := alertCli.AddTeam(addTeamReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", addTeamResponse.Status)
	fmt.Println("code:", addTeamResponse.Code)
}
