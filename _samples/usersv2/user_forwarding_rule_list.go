package main

import (
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/userv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	userCli, _ := cli.UserV2()

	request := userv2.ListUserForwardingRulesRequest{
		Identifier: &userv2.Identifier{UserName: "0"},
	}

	response, err := userCli.ListForwardingRules(request)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, rule := range response.ForwardingRules {
			fmt.Println(rule)
		}
	}
}
