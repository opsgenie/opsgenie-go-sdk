package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

const API_KEY string = "YOUR API KEY HERE"
const PATH_TO_FILE string = "/your/path/to/file/here"

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

	attachFileReq := alerts.AttachFileAlertRequest{AlertId: response.AlertId, Attachment: PATH_TO_FILE, }
	attachFileResp, attachFileErr := alertCli.AttachFile( attachFileReq )
	
	if attachFileErr != nil {
		panic(attachFileErr)
	}

	fmt.Println("Status:", 	attachFileResp.Status)
	fmt.Println("Code:", 	attachFileResp.Code)
}
