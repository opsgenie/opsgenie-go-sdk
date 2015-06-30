package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	policy "github.com/opsgenie/opsgenie-go-sdk/policy"
)

var API_KEY string = "YOUR API KEY HERE"
var POLICY_NAME = "YOUR POLICY NAME HERE"

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)
	cli.SetOpsGenieApiUrl("http://localhost:9000")

	policyCli, cliErr := cli.Policy()

	if cliErr != nil {
		panic(cliErr)
	}
	//disable policy
	disableReq := policy.DisablePolicyRequest{Name: POLICY_NAME}
	_, itgError := policyCli.Disable(disableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Println("Policy disabled successfuly")
	}

	//enable policy
	enableReq := policy.EnablePolicyRequest{Name: POLICY_NAME}
	_, itgError = policyCli.Enable(enableReq)

	if itgError != nil {
		panic(itgError)
	} else {
		fmt.Println("Policy enabled successfuly")
	}
}
