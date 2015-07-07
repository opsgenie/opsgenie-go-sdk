package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
	"github.com/opsgenie/opsgenie-go-sdk/samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(constants.API_KEY)

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

	fmt.Printf("Heartbeat created\n")
	fmt.Printf("-----------------\n")
	fmt.Printf("id: %s\n", response.Id)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// enable the hb
	disableReq := hb.DisableHeartbeatRequest{Id: response.Id}
	disableResp, disableErr := hbCli.Disable(disableReq)
	if disableErr != nil {
		panic(disableErr)
	}

	fmt.Printf("Heartbeat disabled\n")
	fmt.Printf("-----------------\n")
	fmt.Printf("Status: %s\n", disableResp.Status)
	fmt.Printf("Code: %d\n", disableResp.Code)
}
