package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

const API_KEY string = "YOUR API KEY HERE"

const TEAM_NAME string = "YOUR TEAM NAME HERE"

func main() {
	
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)
	
	alertCli, cliErr := cli.Alert()
	
	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandString("Test", 8) }
	response, alertErr := alertCli.Create(req)
	
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:",response.AlertId)
	fmt.Println("status:", 	response.Status)
	fmt.Println("code:", 	response.Code)

	// assign the owner for the alert
	addTeamReq := alerts.AddTeamAlertRequest{ AlertId: response.AlertId, Team: TEAM_NAME, }
	addTeamResponse, alertErr := alertCli.AddTeam(addTeamReq)
	
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", 	addTeamResponse.Status)
	fmt.Println("code:", 	addTeamResponse.Code)
}
