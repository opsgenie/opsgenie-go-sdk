package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/policyv1"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	policyCli, err := cli.PolicyV1()
	if err != nil {
		panic(policyCli)
	}

	response, err := policyCli.List(policyv1.ListPolicyRequest{})
	if err != nil {
		panic(err)
	}

	if len(response.Policies) == 0 {
		fmt.Println("there is no policy")
	} else {
		for _, policy := range response.Policies {
			fmt.Printf(
				"ID: %s\nName:%s\nType: %s\nEnabled: %t\n\n",
				policy.ID,
				policy.Name,
				policy.Type,
				policy.Enabled,
			)
		}
	}
}
