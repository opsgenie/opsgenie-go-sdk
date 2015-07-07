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
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8), Note: "Created for testing purposes"}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertId)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// list alert recipients
	getRecipientsReq := alerts.ListAlertRecipientsRequest{Id: response.AlertId}
	getRecipientsResponse, alertErr := alertCli.ListRecipients(getRecipientsReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("Users:  %v\n", getRecipientsResponse.Users)
	fmt.Printf("Groups:  %v\n", getRecipientsResponse.Groups)
}
