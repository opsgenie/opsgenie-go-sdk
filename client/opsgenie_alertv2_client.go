package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	"github.com/opsgenie/opsgenie-go-sdk/alertsv2/savedsearches"
)

type OpsGenieAlertV2Client struct {
	RestClient
}

func (cli *OpsGenieAlertV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieAlertV2Client) Get(req alertsv2.GetAlertRequest) (*alertsv2.DetailedAlertResponse, error) {
	var response alertsv2.DetailedAlertResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) List(req alertsv2.ListAlertRequest) (*alertsv2.ListAlertResponse, error) {
	var response alertsv2.ListAlertResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) ListAlertRecipients(req alertsv2.ListAlertRecipientsRequest) (*alertsv2.ListAlertRecipientsResponse, error) {
	var response alertsv2.ListAlertRecipientsResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) ListAlertLogs(req alertsv2.ListAlertLogsRequest) (*alertsv2.ListAlertLogsResponse, error) {
	var response alertsv2.ListAlertLogsResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) ListAlertNotes(req alertsv2.ListAlertNotesRequest) (*alertsv2.ListAlertNotesResponse, error) {
	var response alertsv2.ListAlertNotesResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) Acknowledge(req alertsv2.AcknowledgeRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) Close(req alertsv2.CloseRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) Unacknowledge(req alertsv2.UnacknowledgeRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) Snooze(req alertsv2.SnoozeRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) ExecuteCustomAction(req alertsv2.ExecuteCustomActionRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) Create(req alertsv2.CreateAlertRequest) (*AsyncRequestResponse, error) {
	req.Init()
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) Delete(req alertsv2.DeleteAlertRequest) (*AsyncRequestResponse, error) {
	var response AsyncRequestResponse

	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieAlertV2Client) AddNote(req alertsv2.AddNoteRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) EscalateToNext(req alertsv2.EscalateToNextRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) Assign(req alertsv2.AssignAlertRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) AddTeamToAlert(req alertsv2.AddTeamToAlertRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) AddTags(req alertsv2.AddTagsToAlertRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) RemoveTags(req alertsv2.RemoveTagsRequest) (*AsyncRequestResponse, error) {
	var response AsyncRequestResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) AddDetails(req alertsv2.AddDetailsRequest) (*AsyncRequestResponse, error) {
	return cli.sendAsyncPostRequest(&req)
}

func (cli *OpsGenieAlertV2Client) RemoveDetails(req alertsv2.RemoveDetailsRequest) (*AsyncRequestResponse, error)  {
	var response AsyncRequestResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) CreateSavedSearch(req savedsearches.CreateSavedSearchRequest) (*savedsearches.CreateSavedSearchResponse, error) {
	var response savedsearches.CreateSavedSearchResponse

	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) GetSavedSearch(req alertsv2.GetSavedSearchRequest) (*savedsearches.GetSavedSearchResponse, error) {
	var response savedsearches.GetSavedSearchResponse

	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) UpdateSavedSearch(req alertsv2.UpdateSavedSearchRequest) (*savedsearches.UpdateSavedSearchResponse, error) {
	var response savedsearches.UpdateSavedSearchResponse
	err := cli.sendPatchRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieAlertV2Client) ListSavedSearches(req alertsv2.LisSavedSearchRequest) (*savedsearches.ListSavedSearchResponse, error) {
	var response savedsearches.ListSavedSearchResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieAlertV2Client) GetRequestStatus(req alertsv2.GetRequestStatusRequest) (*alertsv2.GetRequestStatusResponse, error) {
	var response alertsv2.GetRequestStatusResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
