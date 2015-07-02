package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

func main() {
	API_KEY := "YOUR API KEY HERE"

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
	enableReq := hb.EnableHeartbeatRequest{Id: response.Id}
	enableResp, enableErr := hbCli.Enable(enableReq)
	if enableErr != nil {
		panic(enableErr)
	}

	fmt.Println("Heartbeat enabled")
	fmt.Println("-----------------")
	fmt.Println("Status:", enableResp.Status)
	fmt.Println("Code:", enableResp.Code)
}
