package main

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
	"github.com/opsgenie/opsgenie-go-sdk/samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(constants.API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test - ", 10), Actions: constants.ACTIONS}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertId)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// execute sample 'pong' action for the alert
	execActionReq := alerts.ExecuteActionAlertRequest{Id: response.AlertId, Action: constants.ACTION_TO_EXEC, Note: "Action <b>pong</b> executed by the Go API"}
	execActionResponse, alertErr := alertCli.ExecuteAction(execActionReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("status: %s\n", execActionResponse.Result)
	fmt.Printf("code: %d\n", execActionResponse.Code)
}
