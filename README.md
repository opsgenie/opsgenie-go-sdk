# opsgenie-go-sdk *beta*

## Aim and Scope
OpsGenie GO SDK aims to access OpsGenie Web API through HTTP calls
from a client application purely written in Go language.

OpsGenie Go SDK covers *the Alert API*, *the Heartbeat API*,
*the Integration API* and *the Policy API*. Future releases
are subject to be delivered for packing more APIs soon.

For more information about OpsGenie Go SDK, please refer to [OpsGenie Go API](https://www.opsgenie.com/docs/api-and-client-libraries/opsgenie-go-api) document.

## Pre-requisites
* The API is built using Go 1.4.1. Some features may not be
available or supported unless you have installed a relevant version of Go.
Please click [https://golang.org/dl/](https://golang.org/dl/) to download and
get more information about installing Go on your computer.
* Make sure you have properly set both `GOROOT` and `GOPATH`
environment variables.

* Before you can begin, you need to sign up [OpsGenie](http://www.opsgenie.com) if you
don't have a valid account yet. Create an API Integration and get your API key.

## Installation
To download all packages in the repo with their dependencies, simply run
`go get github.com/opsgenie/opsgenie-go-sdk/...`.

## Running Tests

OpsGenie Go SDK includes a set of acceptance tests. The acceptance tests will
allow you to examine almost all functionalities of the provided SDK.
Before running the test suit, you should edit the `client_at_test_cfg.yaml`
file and place the correct values of your own. Please keep in mind that it's a
`yaml` file and should conform the YAML file specifications (e.g. no tabs but
 whitespaces).

 `go test` command will run the tests and you will get pass if everything goes
 well.

## Getting Started
There are many useful sample code snippets under `samples` directory for packages.

##Handling Zero value problem with 'omitempty' option in Json

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
about the Web API.

* [Alert API](https://www.opsgenie.com/docs/web-api/alert-api)
* [Heartbeat API](https://www.opsgenie.com/docs/web-api/heartbeat-api)
* [Integration API](https://www.opsgenie.com/docs/web-api/integration-api)
* [Policy API](https://www.opsgenie.com/docs/web-api/policy-api)
