package main

import (
	"fmt"
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

var API_KEY string = "YOUR API KEY HERE"
var USER string = "YOUR USERNAME HERE"

func main() {

	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8), Note: "Created for testing purposes", User: USER}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	// close the alert
	getreq := alerts.GetAlertRequest{Id: response.AlertId}
	getresponse, alertErr := alertCli.Get(getreq)
	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("tags:", getresponse.Tags)
	fmt.Println("count:", getresponse.Count)
	fmt.Println("teams:", getresponse.Teams)
	fmt.Println("recipients:", getresponse.Recipients)
	fmt.Println("tiny id:", getresponse.TinyId)
	fmt.Println("alias:", getresponse.Alias)
	fmt.Println("entity:", getresponse.Entity)
	fmt.Println("id:", getresponse.Id)
	fmt.Println("updated at:", getresponse.UpdatedAt)
	fmt.Println("message:", getresponse.Message)
	fmt.Println("details:", getresponse.Details)
	fmt.Println("source:", getresponse.Source)
	fmt.Println("description:", getresponse.Description)
	fmt.Println("created at:", getresponse.CreatedAt)
	fmt.Println("is seen?:", getresponse.IsSeen)
	fmt.Println("acknowledged?:", getresponse.Acknowledged)
	fmt.Println("owner:", getresponse.Owner)
	fmt.Println("actions:", getresponse.Actions)
	fmt.Println("system data:", getresponse.SystemData)
}
