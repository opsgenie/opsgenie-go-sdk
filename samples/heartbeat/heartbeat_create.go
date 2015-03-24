package main

import (
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

const API_KEY string = "YOUR API KEY HERE"

func main() {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)
	
	alertCli, cliErr := cli.Heartbeat()
	
	if cliErr != nil {
		panic(cliErr)
	}

	// create the hb
	req := hb.AddHeartbeatRequest{Name: samples.RandStringWithPrefix("Test", 4) }
	response, hbErr := alertCli.Add(req)
	
	if hbErr != nil {
		panic(hbErr)
	}

	fmt.Println("id:",		response.Id)
	fmt.Println("status:",	response.Status)
	fmt.Println("code:",	response.Code)
}