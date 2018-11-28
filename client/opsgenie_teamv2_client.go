package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/teamv2"
)

// OpsGenieTeamV2Client is the data type to make User API requests.
type OpsGenieTeamV2Client struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieTeamV2Client.
func (cli *OpsGenieTeamV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// List method retrieves the list of users from OpsGenie.
func (cli *OpsGenieTeamV2Client) List(req teamv2.ListTeamsRequest) (*teamv2.ListTeamsResponse, error) {
	var response teamv2.ListTeamsResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Create method creates new user in OpsGenie.
func (cli *OpsGenieTeamV2Client) Create(req teamv2.CreateTeamRequest) (*teamv2.CreateTeamResponse, error) {
	var response teamv2.CreateTeamResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Get method returns user.
func (cli *OpsGenieTeamV2Client) Get(req teamv2.GetTeamRequest) (*teamv2.GetTeamResponse, error) {
	var response teamv2.GetTeamResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
