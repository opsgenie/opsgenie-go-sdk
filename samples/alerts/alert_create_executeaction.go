package main

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

func main() {
	API_KEY := "YOUR API KEY HERE"
	ACTIONS := []string{"ping", "pong"}
	ACTION_TO_EXEC := "pong"

	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test - ", 10), Actions: ACTIONS}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// execute sample 'pong' action for the alert
	execActionReq := alerts.ExecuteActionAlertRequest{Id: response.AlertId, Action: ACTION_TO_EXEC, Note: "Action <b>pong</b> executed by the Go API"}
	execActionResponse, alertErr := alertCli.ExecuteAction(execActionReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", execActionResponse.Result)
	fmt.Println("code:", execActionResponse.Code)
}
