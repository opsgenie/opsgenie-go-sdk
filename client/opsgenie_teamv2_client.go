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
