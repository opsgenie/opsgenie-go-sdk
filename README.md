# opsgenie-go-sdk

<aside class="notice">
This SDK is not actively maintained. Please consider generating a client from the [OpenAPI Spec](https://github.com/opsgenie/opsgenie-oas)
</aside>

## Aim and Scope
OpsGenie GO SDK aims to access OpsGenie Web API through HTTP calls
from a client application purely written in Go language.

OpsGenie Go SDK covers *the Alert API*, *the Heartbeat API*,
*the Integration API* and *the Policy API*. Future releases
are subject to be delivered for packing more APIs soon.

**Documentation:** [![](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/opsgenie/opsgenie-go-sdk/client)

For more information about OpsGenie Go SDK, please refer to [OpsGenie Go API](https://www.opsgenie.com/docs/api-and-client-libraries/opsgenie-go-api) document.

## Pre-requisites
* The API is built using Go 1.4.2. Some features may not be
available or supported unless you have installed a relevant version of Go.
Please click [https://golang.org/dl/](https://golang.org/dl/) to download and
get more information about installing Go on your computer.
* Make sure you have properly set both `GOROOT` and `GOPATH`
environment variables.

* Before you can begin, you need to sign up [OpsGenie](http://www.opsgenie.com) if you
don't have a valid account yet. Create an API Integration and get your API key.

## Installation
To download all packages in the repo with their dependencies, simply run

`go get github.com/opsgenie/opsgenie-go-sdk/...`

## Getting Started
One can start using OpsGenie Go SDK by initializing client and making a request. Example shown below demonstrates how to initialize an OpsGenie Alert client and make a create alert request.
```go
package main

import (
	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, _ := cli.AlertV2()

	teams := []alertsv2.TeamRecipient{
		&alertsv2.Team{Name: "teamId"},
		&alertsv2.Team{ID: "teamId"},
	}

	visibleTo := []alertsv2.Recipient{
		&alertsv2.Team{ID: "teamId"},
		&alertsv2.Team{Name: "teamName"},
		&alertsv2.User{ID: "userId"},
		&alertsv2.User{Username: "user@opsgenie.com"},
	}

	request := alertsv2.CreateAlertRequest{
			Message:     "message",
			Alias:       "alias",
			Description: "alert description",
			Teams:       teams,
			VisibleTo:   visibleTo,
			Actions:     []string{"action1", "action2"},
			Tags:        []string{"tag1", "tag2"},
			Details: map[string]string{
				"key":  "value",
				"key2": "value2",
			},
			Entity:   "entity",
			Source:   "source",
			Priority: alertsv2.P1,
			User:     "user@opsgenie.com",
			Note:     "alert note",
		}

	response, err := alertCli.Create(request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Create request ID: " + response.RequestID)
	}
}
    }
```
There are many useful sample code snippets under `samples` directory for packages.

## Handling Zero value problem with 'omitempty' option in Json

Every golang type has a [zero value](http://golang.org/ref/spec#The_zero_value).
AddHeartbeat and UpdateHeartbeat requests have a boolean "Enabled" field to determine a heartbeat is enabled or disabled.
enabled is not a mandatory field in both requests so, it has "omitempty" flag.

When JSON is unmarshalling the requests, it omits the fields contains zero value of its type when this option is set.

The problem starts here:
When you want to set the Enabled field as false, JSON does not unmarshal it because it contains boolean zero value and has option omitempty.
This problem occurs with strings; when you want to reset a heartbeat's description. To set heartbeat's description as empty string, you should make the request with Description :"".
But JSON does not unmarshal it either.

So, to solve this we followed go-github's solution as mentioned [here](https://willnorris.com/2014/05/go-rest-apis-and-pointers).
We used pointers just for booleans, if you want to set a string's value to empty. please use " ", or "-" as new string value.

## The Web API

Please follow the links below for more information and details
about the REST API.

* [Alert API](https://www.opsgenie.com/docs/rest-api/alert-api)
* [Heartbeat API](https://www.opsgenie.com/docs/rest-api/heartbeat-api)
* [Integration API](https://www.opsgenie.com/docs/rest-api/integration-api)
* [Policy API](hhttps://www.opsgenie.com/docs/rest-api/policy-api)


## Bug Reporting and Feature Requests

If you like to report a bug, or a feature request; please open an issue.
