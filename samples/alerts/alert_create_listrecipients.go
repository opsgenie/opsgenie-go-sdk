package main

import (
	"fmt"
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8), Note: "Created for testing purposes"}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// list alert recipients
	getRecipientsReq := alerts.ListAlertRecipientsRequest{Id: response.AlertId}
	getRecipientsResponse, alertErr := alertCli.ListRecipients(getRecipientsReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("Users: ", getRecipientsResponse.Users)
	fmt.Println("Groups: ", getRecipientsResponse.Groups)
}
