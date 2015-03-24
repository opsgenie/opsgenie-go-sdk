# opsgenie-go-sdk *beta*

## Aim and Scope
OpsGenie GO SDK aims to access OpsGenie Web API through HTTP calls 
from a client application purely written in Go language.

Beta release covers *the Alert API* and *the Heartbeat API*. Future releases 
are subject to be delivered for packing more APIs soon.

## Pre-requisites
* The API is built using Go 1.4.1. Some features may not be 
available or supported unless you have installed a relevant version of Go. 
Please click [https://golang.org/dl/](https://golang.org/dl/) to download and 
get more information about installing Go on your computer.
* Make sure you have properly set both `GOROOT` and `GOPATH` 
environment variables.
* Obviously, you need to sign up [OpsGenie](http://www.opsgenie.com) if you 
don't have a valid account yet. You must grab two separate api keys, one for 
accessing the Alerts API and one for the Heartbeat API. 


## Installation

Simply run `go get https://github.com/opsgenie/opsgenie-go-sdk`.

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

There are many useful sample code snippets under `samples` directory for both 
*alerts* and *heartbeat* packages.

## The Web API

Please follow the links below for more information and details 
about the Web API.

* [Alert API](https://www.opsgenie.com/docs/web-api/alert-api)
* [Heartbeat API](https://www.opsgenie.com/docs/web-api/heartbeat-api)
