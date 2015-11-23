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
	cli.SetAPIKey(constants.APIKey)

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
	fmt.Printf("name: %s\n", response.Name)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// enable the hb
	getReq := hb.GetHeartbeatRequest{Name: response.Name}
	getResp, getErr := hbCli.Get(getReq)
	if getErr != nil {
		panic(getErr)
	}

	fmt.Printf("Heartbeat details\n")
	fmt.Printf("-----------------\n")
	fmt.Printf("Name: %s\n", getResp.Name)
	fmt.Printf("Status: %s\n", getResp.Status)
	fmt.Printf("Description: %s\n", getResp.Description)
	fmt.Printf("Enabled?: %t\n", getResp.Enabled)
	fmt.Printf("Last Heartbeat: %d\n", getResp.LastHeartbeat)
	fmt.Printf("Interval: %d\n", getResp.Interval)
	fmt.Printf("Interval Unit: %s\n", getResp.IntervalUnit)
	fmt.Printf("Expired?: %t\n", getResp.Expired)
}
