package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/contactv2"
)

type OpsGenieContactV2Client struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieAlertV2Client.
func (cli *OpsGenieContactV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieContactV2Client) Create(req contactv2.CreateContactRequest) (*contactv2.CreateContactResponse, error) {
	var response contactv2.CreateContactResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
