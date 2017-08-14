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

	request := userv2.ListUserEscalationsRequest{
		Identifier: &userv2.Identifier{UserName: "0"},
	}

	response, err := userCli.ListEscalations(request)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, escalation := range response.Escalations {
			fmt.Println(escalation)
		}
	}
}
