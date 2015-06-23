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

	fmt.Println("Heartbeat created")
	fmt.Println("-----------------")
	fmt.Println("id:",		response.Id)
	fmt.Println("status:",	response.Status)
	fmt.Println("code:",	response.Code)

	// enable the hb
	getReq := hb.GetHeartbeatRequest{Id: response.Id}
	getResp, getErr := hbCli.Get( getReq )
	if getErr != nil {
		panic(getErr)
	}

	fmt.Println("Heartbeat details")
	fmt.Println("-----------------")
	fmt.Println("Id:",		getResp.Id) 
	fmt.Println("Name:",	getResp.Name)
	fmt.Println("Status",	getResp.Status)
	fmt.Println("Description:", getResp.Description)
	fmt.Println("Enabled?:", getResp.Enabled)
	fmt.Println("Last Heart beat:", getResp.LastHeartbeat)
	fmt.Println("Interval:", getResp.Interval)
	fmt.Println("Interval Unit:", getResp.IntervalUnit)
	fmt.Println("Expired?:", getResp.Expired)
}