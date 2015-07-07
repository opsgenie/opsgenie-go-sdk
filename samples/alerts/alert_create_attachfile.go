package main

import (
	"fmt"
	"os"

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
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Printf("message: %s\n", response.Message)
	fmt.Printf("alert id: %s\n", response.AlertId)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	file, err := os.OpenFile(constants.PATH_TO_FILE, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	attachFileReq := alerts.AttachFileAlertRequest{Id: response.AlertId, Attachment: file}
	attachFileResp, attachFileErr := alertCli.AttachFile(attachFileReq)

	if attachFileErr != nil {
		panic(attachFileErr)
	}

	fmt.Printf("Status: %s\n", attachFileResp.Status)
	fmt.Printf("Code: %d\n", attachFileResp.Code)
}
