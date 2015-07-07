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

	hbName := samples.RandStringWithPrefix("Test", 4)
	// create the hb
	req := hb.AddHeartbeatRequest{Name: hbName}
	response, hbErr := hbCli.Add(req)

	if hbErr != nil {
		panic(hbErr)
	}

	fmt.Printf("Heartbeat added\n")
	fmt.Printf("---------------\n")
	fmt.Printf("id: %s\n", response.Id)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// send heart beat request
	sendReq := hb.SendHeartbeatRequest{Name: hbName}
	sendResp, sendErr := hbCli.Send(sendReq)

	if sendErr != nil {
		panic(sendErr)
	}

	fmt.Printf("Heartbeat request sent\n")
	fmt.Printf("----------------------\n")
	fmt.Printf("Heartbeat: %d\n", sendResp.Heartbeat)
	fmt.Printf("Will expire at: %d\n", sendResp.WillExpireAt)
	fmt.Printf("Status: %s\n", sendResp.Status)
	fmt.Printf("Code: %d\n", sendResp.Code)
}
