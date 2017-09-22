package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/policyv1"
)

type OpsGeniePolicyV1Client struct {
	RestClient
}

func (cli *OpsGeniePolicyV1Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGeniePolicyV1Client) CreateAutoClosePolicy(req policyv1.CreateAutoClosePolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	return cli.create(&req)
}

func (cli *OpsGeniePolicyV1Client) CreateAutoRestartPolicy(req policyv1.CreateAutoRestartPolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	return cli.create(&req)
}

func (cli *OpsGeniePolicyV1Client) CreateNotificationSuppressPolicy(req policyv1.CreateNotificationSuppressPolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	return cli.create(&req)
}

func (cli *OpsGeniePolicyV1Client) CreateNotificationDeduplicationPolicy(req policyv1.CreateNotificationDeduplicationPolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	return cli.create(&req)
}

func (cli *OpsGeniePolicyV1Client) CreateNotificationDelayPolicy(req policyv1.CreateNotificationDelayPolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	return cli.create(&req)
}

func (cli *OpsGeniePolicyV1Client) CreateModifyPolicy(req policyv1.CreateModifyPolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	return cli.create(&req)
}

func (cli *OpsGeniePolicyV1Client) create(req Request) (*policyv1.CreatePolicyResponse, error) {
	var response policyv1.CreatePolicyResponse
	err := cli.sendPostRequest(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Get(req policyv1.GetPolicyRequest) (*policyv1.GetPolicyResponse, error) {
	var response policyv1.GetPolicyResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) UpdateAutoClosePolicy(req policyv1.UpdateAutoClosePolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	return cli.update(&req)
}

func (cli *OpsGeniePolicyV1Client) UpdateAutoRestartPolicy(req policyv1.UpdateAutoRestartPolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	return cli.update(&req)
}

func (cli *OpsGeniePolicyV1Client) UpdateNotificationSuppressPolicy(req policyv1.UpdateNotificationSuppressPolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	return cli.update(&req)
}

func (cli *OpsGeniePolicyV1Client) UpdateNotificationDeduplicationPolicy(req policyv1.UpdateNotificationDeduplicationPolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	return cli.update(&req)
}

func (cli *OpsGeniePolicyV1Client) UpdateNotificationDelayPolicy(req policyv1.UpdateNotificationDelayPolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	return cli.update(&req)
}

func (cli *OpsGeniePolicyV1Client) UpdateModifyPolicy(req policyv1.UpdateModifyPolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	return cli.update(&req)
}

func (cli *OpsGeniePolicyV1Client) update(req Request) (*policyv1.UpdatePolicyResponse, error) {
	var response policyv1.UpdatePolicyResponse
	err := cli.sendPutRequest(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Delete(req policyv1.DeletePolicyRequest) (*policyv1.DeletePolicyResponse, error) {
	var response policyv1.DeletePolicyResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Disable(req policyv1.DisablePolicyRequest) (*policyv1.DisablePolicyResponse, error) {
	var response policyv1.DisablePolicyResponse
	req.StatusAction = policyv1.DisableAction
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Enable(req policyv1.EnablePolicyRequest) (*policyv1.EnablePolicyResponse, error) {
	var response policyv1.EnablePolicyResponse
	req.StatusAction = policyv1.EnableAction
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) List(req policyv1.ListPolicyRequest) (*policyv1.ListPolicyResponse, error) {
	var response policyv1.ListPolicyResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
