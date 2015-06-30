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
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// add recipient
	addRecipientReq := alerts.AddRecipientAlertRequest{Id: response.AlertId, Recipient: "recipient"}
	addRecipientResponse, alertErr := alertCli.AddRecipient(addRecipientReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", addRecipientResponse.Status)
	fmt.Println("code:", addRecipientResponse.Code)
}
