package main

import (
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"

func main() {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)
	
	hbCli, cliErr := cli.Heartbeat()
	
	if cliErr != nil {
		panic(cliErr)
	}

	// create the hb
	req := hb.AddHeartbeatRequest{Name: samples.RandStringWithPrefix("Test", 4) }
	response, hbErr := hbCli.Add(req)
	
	if hbErr != nil {
		panic(hbErr)
	}

	fmt.Println("Heartbeat added")
	fmt.Println("---------------")
	fmt.Println("id:",		response.Id)
	fmt.Println("status:",	response.Status)
	fmt.Println("code:",	response.Code)

	// update the newly created heart beat, change the name
	hbId := response.Id
	updateReq := hb.UpdateHeartbeatRequest{Id: hbId, Name: samples.RandStringWithPrefix("Updated Test", 8), Description: response.Id + " is getting updated" }
	updateResp, updateErr := hbCli.Update(updateReq)

	if updateErr != nil {
		panic(updateErr)
	}

	fmt.Println("Heartbeat updated")
	fmt.Println("-----------------")
	fmt.Println("id:",		updateResp.Id)
	fmt.Println("status:",	updateResp.Status)
	fmt.Println("code:",	updateResp.Code)	
}