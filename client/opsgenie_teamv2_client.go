package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/teamv2"
)

type OpsGenieTeamV2Client struct {
	RestClient
}

func (cli *OpsGenieTeamV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieTeamV2Client) Create(req teamv2.CreateTeamRequest) (*teamv2.CreateTeamResponse, error) {
	var response teamv2.CreateTeamResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieTeamV2Client) Get(req teamv2.GetTeamRequest) (*teamv2.GetTeamResponse, error) {
	var response teamv2.GetTeamResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
