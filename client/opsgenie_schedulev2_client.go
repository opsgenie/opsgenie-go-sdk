package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/schedulev2"
)

type OpsGenieScheduleV2Client struct {
	RestClient
}

func (cli *OpsGenieScheduleV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieScheduleV2Client) Create(req schedulev2.CreateScheduleRequest) (*schedulev2.CreateScheduleResponse, error) {
	var response schedulev2.CreateScheduleResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieScheduleV2Client) Get(req schedulev2.GetScheduleRequest) (*schedulev2.GetScheduleResponse, error) {
	var response schedulev2.GetScheduleResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
