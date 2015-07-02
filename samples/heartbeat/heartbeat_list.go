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

	// list the HBs
	listReq := hb.ListHeartbeatsRequest{}
	listResp, listErr := hbCli.List(listReq)
	if listErr != nil {
		panic(listErr)
	}

	fmt.Println("Heartbeats")
	fmt.Println("-----------------")
	beats := listResp.Heartbeats
	for _, beat := range beats {
		fmt.Println("Id:", beat.Id)
		fmt.Println("Name:", beat.Name)
		fmt.Println("Status", beat.Status)
		fmt.Println("Description:", beat.Description)
		fmt.Println("Enabled?:", beat.Enabled)
		fmt.Println("Last Heart beat:", beat.LastHeartbeat)
		fmt.Println("Interval:", beat.Interval)
		fmt.Println("Interval Unit:", beat.IntervalUnit)
		fmt.Println("Expired?:", beat.Expired)
		fmt.Println("-----------------")
	}
}
