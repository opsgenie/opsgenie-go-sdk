package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/policy"
	"github.com/opsgenie/opsgenie-go-sdk/samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(constants.API_KEY)

	policyCli, cliErr := cli.Policy()

	if cliErr != nil {
		panic(cliErr)
	}
	//disable policy
	disableReq := policy.DisablePolicyRequest{Name: constants.POLICY_NAME}
	_, itgError := policyCli.Disable(disableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Printf("Policy disabled successfuly\n")
	}

	//enable policy
	enableReq := policy.EnablePolicyRequest{Name: constants.POLICY_NAME}
	_, itgError = policyCli.Enable(enableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Printf("Policy enabled successfuly\n")
	}
}
