package main

import (
	itg "github.com/opsgenie/opsgenie-go-sdk/integration"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
)

var API_KEY string = "YOUR API KEY HERE"
var INTEGRATION_NAME = "YOUR INTEGRATION NAME HERE"
func main() {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)
	cli.SetOpsGenieApiUrl("http://localhost:9000")

	integrationCli, cliErr := cli.Integration()
	
	if cliErr != nil {
		panic(cliErr)
	}
	//disable integration
	disableReq := itg.DisableIntegrationRequest{Name : INTEGRATION_NAME}
	_, itgError := integrationCli.Disable(disableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Println("Integration disabled successfuly")
	}

	//enable integration
	enableReq := itg.EnableIntegrationRequest{Name: INTEGRATION_NAME}
	_, itgError = integrationCli.Enable(enableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Println("Integration enabled successfuly")
	}
}