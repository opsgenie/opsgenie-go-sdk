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

	request := userv2.CreateUserRequest{
		UserName:      "user-name@test.com",
		FullName:      "User Name",
		Role:          userv2.Role{Name: userv2.UserRole},
		SkypeUsername: "user.name",
		UserAddress: userv2.UserAddress{
			Country: "US",
			State:   "Indiana",
			City:    "Terre Haute",
			Line:    "567 Stratford Park",
			ZipCode: "47802",
		},
		Tags:               userv2.Tags{"advanced", "marked"},
		Details:            userv2.Details{"detail1key": []string{"detail1dvalue1", "detail1value2"}},
		Timezone:           "Europe/Moscow",
		Locale:             "en_US",
		InvitationDisabled: false,
	}

	response, err := userCli.Create(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Id: %s, username: %#v\n", response.Data.ID, response.Data.Name)
	}
}
