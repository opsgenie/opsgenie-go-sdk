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
	deleteReq := hb.DeleteHeartbeatRequest{Id: response.Id}
	deleteResp, deleteErr := hbCli.Delete(deleteReq)
	if deleteErr != nil {
		panic(deleteErr)
	}

	fmt.Println("Heartbeat deleted")
	fmt.Println("-----------------")
	fmt.Println("Status:", deleteResp.Status)
	fmt.Println("Code:", deleteResp.Code)
}
