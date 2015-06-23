package client

import (
	integration "github.com/opsgenie/opsgenie-go-sdk/integration"
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
)

const (
	ENABLE_INTEGRATION_URL 		= "/v1/json/integration/enable"
	DISABLE_INTEGRATION_URL 	= "/v1/json/integration/disable"
)

type OpsGenieIntegrationClient struct {
	OpsGenieClient
}

func (cli *OpsGenieIntegrationClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieIntegrationClient) Enable(req integration.EnableIntegrationRequest) (*integration.EnableIntegrationResponse, error){
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ENABLE_INTEGRATION_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var enableIntegrationResp integration.EnableIntegrationResponse
		if err = resp.Body.FromJsonTo(&enableIntegrationResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &enableIntegrationResp, nil
	}
}

func (cli *OpsGenieIntegrationClient) Disable(req integration.DisableIntegrationRequest) (*integration.DisableIntegrationResponse, error){
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(DISABLE_INTEGRATION_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var disableIntegrationResp integration.DisableIntegrationResponse
		if err = resp.Body.FromJsonTo(&disableIntegrationResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &disableIntegrationResp, nil
	}
}
