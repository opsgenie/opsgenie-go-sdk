package client

import (
	"errors"

	"github.com/opsgenie/opsgenie-go-sdk/group"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
)

const (
	groupURL          = "/v1/json/group"
)

// OpsGenieGroupClient is the data type to make Group API requests.
type OpsGenieGroupClient struct {
	OpsGenieClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieGroupClient.
func (cli *OpsGenieGroupClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// Create method creates a group at OpsGenie.
func (cli *OpsGenieGroupClient) Create(req group.CreateGroupRequest) (*group.CreateGroupResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(groupURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createGroupResp group.CreateGroupResponse

	if err = resp.Body.FromJsonTo(&createGroupResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &createGroupResp, nil
}

// Update method updates a group at OpsGenie.
func (cli *OpsGenieGroupClient) Update(req group.UpdateGroupRequest) (*group.UpdateGroupResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(groupURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updateGroupResp group.UpdateGroupResponse

	if err = resp.Body.FromJsonTo(&updateGroupResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &updateGroupResp, nil
}

// Delete method deletes a group at OpsGenie.
func (cli *OpsGenieGroupClient) Delete(req group.DeleteGroupRequest) (*group.DeleteGroupResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildDeleteRequest(groupURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deleteGroupResp group.DeleteGroupResponse

	if err = resp.Body.FromJsonTo(&deleteGroupResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &deleteGroupResp, nil
}

// Get method retrieves specified group details from OpsGenie.
func (cli *OpsGenieGroupClient) Get(req group.GetGroupRequest) (*group.GetGroupResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(groupURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getGroupResp group.GetGroupResponse

	if err = resp.Body.FromJsonTo(&getGroupResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &getGroupResp, nil
}

// List method retrieves groups from OpsGenie.
func (cli *OpsGenieGroupClient) List(req group.ListGroupsRequest) (*group.ListGroupsResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(groupURL,req))
	if resp == nil {
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()

	var listGroupsResp group.ListGroupsResponse

	if err = resp.Body.FromJsonTo(&listGroupsResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &listGroupsResp, nil
}