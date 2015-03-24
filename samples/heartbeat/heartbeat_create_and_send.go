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
	
	hbCli, cliErr := cli.Heartbeat()
	
	if cliErr != nil {
		panic(cliErr)
	}

	hbName := samples.RandStringWithPrefix("Test", 4)
	// create the hb
	req := hb.AddHeartbeatRequest{Name: hbName }
	response, hbErr := hbCli.Add(req)
	
	if hbErr != nil {
		panic(hbErr)
	}

	fmt.Println("Heartbeat added")
	fmt.Println("---------------")
	fmt.Println("id:",		response.Id)
	fmt.Println("status:",	response.Status)
	fmt.Println("code:",	response.Code)

	// send heart beat request
	sendReq := hb.SendHeartbeatRequest{Name: hbName}
	sendResp, sendErr := hbCli.Send(sendReq)

	if sendErr != nil {
		panic(sendErr)
	}

	fmt.Println("Heartbeat request sent")
	fmt.Println("----------------------")
	fmt.Println("Heart beat:",		sendResp.Heartbeat)
	fmt.Println("Will expire at:",	sendResp.WillExpireAt)
	fmt.Println("Status:",			sendResp.Status)
	fmt.Println("Code:", 			sendResp.Code)	
}