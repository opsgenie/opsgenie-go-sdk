package main

import (
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	notificationCli, _ := cli.NotificationV2()

	response, err := notificationCli.Delete(notificationv2.DeleteNotificationRequest{
		Identifier: &notificationv2.Identifier{
			UserID: "0",
			RuleID: "0",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Result: %s, Request ID: %s\n", response.Result, response.RequestID)
	}
}
