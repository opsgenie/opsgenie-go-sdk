package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/escalationv2"
)

type OpsGenieEscalationV2Client struct {
	RestClient
}

func (cli *OpsGenieEscalationV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieEscalationV2Client) List(req escalationv2.ListEscalationRequest) (*escalationv2.ListEscalationResponse, error) {
	var response escalationv2.ListEscalationResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
