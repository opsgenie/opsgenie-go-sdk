package main

import (
	"fmt"
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"
var OWNER string = "YOUR USERNAME HERE"

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

	// assign the owner for the alert
	assignOwnerReq := alerts.AssignOwnerAlertRequest{Id: response.AlertId, Owner: OWNER}
	assignOwnerResponse, alertErr := alertCli.AssignOwner(assignOwnerReq)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", assignOwnerResponse.Status)
	fmt.Println("code:", assignOwnerResponse.Code)
}
