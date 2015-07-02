package main

import (
	"fmt"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	samples "github.com/opsgenie/opsgenie-go-sdk/samples"
)

func main() {
	API_KEY := "YOUR API KEY HERE"

	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// create the alert
	req := alerts.CreateAlertRequest{Message: samples.RandStringWithPrefix("Test", 8)}
	response, alertErr := alertCli.Create(req)

	if alertErr != nil {
		panic(alertErr)
	}

	fmt.Println("message:", response.Message)
	fmt.Println("alert id:", response.AlertId)
	fmt.Println("status:", response.Status)
	fmt.Println("code:", response.Code)

	addnotereq := alerts.AddNoteAlertRequest{}
	// add alert ten notes
	for i := 0; i < 10; i++ {
		addnotereq.Id = response.AlertId
		addnotereq.Note = samples.RandString(45)
		addnoteresp, alertErr := alertCli.AddNote(addnotereq)
		if alertErr != nil {
			panic(alertErr)
		}
		fmt.Println("[Add note] ", addnoteresp.Status, addnoteresp.Code)
	}
	listNotesReq := alerts.ListAlertNotesRequest{Id: response.AlertId}
	listNotesResponse, alertErr := alertCli.ListNotes(listNotesReq)
	if alertErr != nil {
		panic(alertErr)
	}

	alertNotes := listNotesResponse.Notes

	fmt.Println("Last key:", listNotesResponse.LastKey)
	fmt.Println("Notes:")
	fmt.Println("------")

	for _, note := range alertNotes {
		fmt.Println("Note:", note.Note)
		fmt.Println("Owner:", note.Owner)
		fmt.Println("Created at:", note.CreatedAt)
		fmt.Println("-------------------------")
	}
}
