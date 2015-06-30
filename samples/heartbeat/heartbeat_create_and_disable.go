package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	hbCli, cliErr := cli.Heartbeat()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the hb
	req := hb.AddHeartbeatRequest{Name: samples.RandStringWithPrefix("Test", 4)}
	response, hbErr := hbCli.Add(req)

	if hbErr != nil {
		panic(hbErr)
	}

	fmt.Println("Heartbeat created")
	fmt.Println("-----------------")
	fmt.Println("id:", response.Id)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// enable the hb
	disableReq := hb.DisableHeartbeatRequest{Id: response.Id}
	disableResp, disableErr := hbCli.Disable(disableReq)
	if disableErr != nil {
		panic(disableErr)
	}

	fmt.Println("Heartbeat disabled")
	fmt.Println("-----------------")
	fmt.Println("Status:", disableResp.Status)
	fmt.Println("Code:", disableResp.Code)
}
