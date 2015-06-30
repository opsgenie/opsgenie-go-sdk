package main

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"
var USER string = "YOUR USERNAME HERE"

func main() {

	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8), Note: "Created for testing purposes", User: USER}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// close the alert
	getLogsReq := alerts.ListAlertLogsRequest{Id: response.AlertId}
	getLogsResponse, alertErr := alertCli.ListLogs(getLogsReq)
	if alertErr != nil {
		panic(alertErr)
	}

	logs := getLogsResponse.Logs
	for _, log := range logs {
		fmt.Println("Owner:", log.Owner)
		fmt.Println("Log:", log.Log)
		fmt.Println("Log type:", log.LogType)
		fmt.Println("Created at:", log.CreatedAt)
	}
}
