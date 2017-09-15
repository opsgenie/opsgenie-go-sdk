package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/policyv1"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	policyCli, err := cli.PolicyV1()
	if err != nil {
		panic(policyCli)
	}

	response, err := policyCli.Enable(policyv1.EnablePolicyRequest{
		Identifier: &policyv1.Identifier{ID:"1"},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", response.Status)
}
