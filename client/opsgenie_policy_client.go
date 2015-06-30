package client

import (
	"errors"

	"github.com/opsgenie/opsgenie-go-sdk/logging"
	policy "github.com/opsgenie/opsgenie-go-sdk/policy"
)

const (
	ENABLE_POLICY_URL  = "/v1/json/alert/policy/enable"
	DISABLE_POLICY_URL = "/v1/json/alert/policy/disable"
)

type OpsGeniePolicyClient struct {
	OpsGenieClient
}

func (cli *OpsGeniePolicyClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGeniePolicyClient) Enable(req policy.EnablePolicyRequest) (*policy.EnablePolicyResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ENABLE_POLICY_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var enablePolicyResp policy.EnablePolicyResponse
		if err = resp.Body.FromJsonTo(&enablePolicyResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &enablePolicyResp, nil
	}
}

func (cli *OpsGeniePolicyClient) Disable(req policy.DisablePolicyRequest) (*policy.DisablePolicyResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(DISABLE_POLICY_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var disablePolicyResp policy.DisablePolicyResponse
		if err = resp.Body.FromJsonTo(&disablePolicyResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &disablePolicyResp, nil
	}
}
