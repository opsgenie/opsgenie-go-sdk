package main

import (
	"encoding/json"
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	notificationCli, _ := cli.NotificationV2()

	response, err := notificationCli.Get(notificationv2.GetNotificationRequest{
		Identifier: &notificationv2.Identifier{
			UserID: "0",
			RuleID: "0",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		notification := response.Notification
		notificationJson, _ := json.Marshal(notification)

		fmt.Println(string(notificationJson))
	}
}
