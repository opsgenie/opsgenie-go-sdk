package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

const API_KEY string = "YOUR API KEY HERE"

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

	// take ownership of the alert
	takeOwnershipReq := alerts.TakeOwnershipAlertRequest{ AlertId: response.AlertId }
	takeOwnershipResponse, alertErr := alertCli.TakeOwnership(takeOwnershipReq)
	
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", 	takeOwnershipResponse.Status)
	fmt.Println("code:", 	takeOwnershipResponse.Code)
}
