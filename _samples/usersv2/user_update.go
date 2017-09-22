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

	request := userv2.UpdateUserRequest{
		Identifier: &userv2.Identifier{
			ID:     "0",
			Expand: userv2.ContactExpandableField,
		},
		FullName:      "Lex Luthor",
		Role:          &userv2.Role{Name: userv2.AdminRole},
		SkypeUsername: "lex.luthor",
		Tags:          &userv2.Tags{"updated"},
		Details:       &userv2.Details{"test": []string{"updated"}},
		Locale:        "de_CH",
		Timezone:      "US/Arizona",
		UserAddress: &userv2.UserAddress{
			City:  "Phoenix",
			State: "Arizona",
		},
	}

	response, err := userCli.Update(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Result)
	}
}
