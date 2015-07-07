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

	fmt.Printf("Heartbeat added\n")
	fmt.Printf("---------------\n")
	fmt.Printf("id: %s\n", response.Id)
	fmt.Printf("status: %s\n", response.Status)
	fmt.Printf("code: %d\n", response.Code)

	// update the newly created heart beat, change the name
	hbId := response.Id
	updateReq := hb.UpdateHeartbeatRequest{Id: hbId, Name: samples.RandStringWithPrefix("Updated Test", 8), Description: response.Id + " is getting updated"}
	updateResp, updateErr := hbCli.Update(updateReq)

	if updateErr != nil {
		panic(updateErr)
	}

	fmt.Printf("Heartbeat updated\n")
	fmt.Printf("-----------------\n")
	fmt.Printf("id: %s\n", updateResp.Id)
	fmt.Printf("status: %s\n", updateResp.Status)
	fmt.Printf("code: %d\n", updateResp.Code)
}
