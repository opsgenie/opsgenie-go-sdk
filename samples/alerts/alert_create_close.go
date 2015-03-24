package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

const API_KEY string = "YOUR API KEY HERE"

var NOTIFY_ARR []string = []string{"YOUR USERNAME HERE"}

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
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// close the alert
	cloreq := alerts.CloseAlertRequest{AlertId: response.AlertId, Notify: NOTIFY_ARR }
	cloresponse, alertErr := alertCli.Close(cloreq)
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("status:", cloresponse.Status)
	fmt.Println("code:", cloresponse.Code)
}
