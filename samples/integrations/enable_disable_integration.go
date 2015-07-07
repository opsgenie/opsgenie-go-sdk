package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	itg "github.com/opsgenie/opsgenie-go-sdk/integration"
	"github.com/opsgenie/opsgenie-go-sdk/samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(constants.API_KEY)
	cli.SetOpsGenieApiUrl("http://localhost:9000")

	integrationCli, cliErr := cli.Integration()

	if cliErr != nil {
		panic(cliErr)
	}
	//disable integration
	disableReq := itg.DisableIntegrationRequest{Name: constants.INTEGRATION_NAME}
	_, itgError := integrationCli.Disable(disableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Printf("Integration disabled successfuly\n")
	}

	//enable integration
	enableReq := itg.EnableIntegrationRequest{Name: constants.INTEGRATION_NAME}
	_, itgError = integrationCli.Enable(enableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Printf("Integration enabled successfuly\n")
	}
}
