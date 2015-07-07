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
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8), Note: "Created for testing purposes", User: constants.USER}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertId)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// close the alert
	getLogsReq := alerts.ListAlertLogsRequest{Id: response.AlertId}
	getLogsResponse, alertErr := alertCli.ListLogs(getLogsReq)
	if alertErr != nil {
		panic(alertErr)
	}

	logs := getLogsResponse.Logs
	for _, log := range logs {
		fmt.Printf("Owner: %s\n", log.Owner)
		fmt.Printf("Log: %s\n", log.Log)
		fmt.Printf("Log type: %s\n", log.LogType)
		fmt.Printf("Created at: %d\n", log.CreatedAt)
	}
}
