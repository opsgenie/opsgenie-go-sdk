package client

import (
	"errors"

	"github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
)

const (
	ADD_HEARTBEAT_URL     = "/v1/json/heartbeat"
	UPDATE_HEARTBEAT_URL  = "/v1/json/heartbeat"
	ENABLE_HEARTBEAT_URL  = "/v1/json/heartbeat/enable"
	DISABLE_HEARTBEAT_URL = "/v1/json/heartbeat/disable"
	DELETE_HEARTBEAT_URL  = "/v1/json/heartbeat"
	GET_HEARTBEAT_URL     = "/v1/json/heartbeat"
	LIST_HEARTBEAT_URL    = "/v1/json/heartbeat"
	SEND_HEARTBEAT_URL    = "/v1/json/heartbeat/send"
)

type OpsGenieHeartbeatClient struct {
	OpsGenieClient
}

func (cli *OpsGenieHeartbeatClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieHeartbeatClient) Add(req heartbeat.AddHeartbeatRequest) (*heartbeat.AddHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey

	resp, err := cli.sendRequest(cli.buildPostRequest(ADD_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var addHeartbeatResp heartbeat.AddHeartbeatResponse
		if err = resp.Body.FromJsonTo(&addHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &addHeartbeatResp, nil
	}
}

// Update Heartbeat is used to change configuration of existing heartbeats.
// Mandatory Parameters:
// 	- id: 		Id of the heartbeat
// 	- apiKey: 	API key is used for authenticating API requests
// Optional Parameters
// 	- name: 			Name of the heartbeat
// 	- interval: 		Specifies how often a heartbeat message should be expected.
// 	- intervalUnit: 	interval specified as minutes, hours or days
// 	- description:	 	An optional description of the heartbeat
// 	- enabled: 			Enable/disable heartbeat monitoring
func (cli *OpsGenieHeartbeatClient) Update(req heartbeat.UpdateHeartbeatRequest) (*heartbeat.UpdateHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(UPDATE_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var updateHeartbeatResp heartbeat.UpdateHeartbeatResponse
		if err = resp.Body.FromJsonTo(&updateHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &updateHeartbeatResp, nil
	}
}

func (cli *OpsGenieHeartbeatClient) Enable(req heartbeat.EnableHeartbeatRequest) (*heartbeat.EnableHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ENABLE_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var enableHeartbeatResp heartbeat.EnableHeartbeatResponse
		if err = resp.Body.FromJsonTo(&enableHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &enableHeartbeatResp, nil
	}
}

func (cli *OpsGenieHeartbeatClient) Disable(req heartbeat.DisableHeartbeatRequest) (*heartbeat.DisableHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(DISABLE_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var disableHeartbeatResp heartbeat.DisableHeartbeatResponse
		if err = resp.Body.FromJsonTo(&disableHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &disableHeartbeatResp, nil
	}

}

func (cli *OpsGenieHeartbeatClient) Delete(req heartbeat.DeleteHeartbeatRequest) (*heartbeat.DeleteHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildDeleteRequest(DELETE_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var deleteHeartbeatResp heartbeat.DeleteHeartbeatResponse
		if err = resp.Body.FromJsonTo(&deleteHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &deleteHeartbeatResp, nil
	}
}

func (cli *OpsGenieHeartbeatClient) Get(req heartbeat.GetHeartbeatRequest) (*heartbeat.GetHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(GET_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var getHeartbeatResp heartbeat.GetHeartbeatResponse
		if err = resp.Body.FromJsonTo(&getHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &getHeartbeatResp, nil
	}
}

func (cli *OpsGenieHeartbeatClient) List(req heartbeat.ListHeartbeatsRequest) (*heartbeat.ListHeartbeatsResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(LIST_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var listHeartbeatsResp heartbeat.ListHeartbeatsResponse
		if err = resp.Body.FromJsonTo(&listHeartbeatsResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &listHeartbeatsResp, nil
	}
}

func (cli *OpsGenieHeartbeatClient) Send(req heartbeat.SendHeartbeatRequest) (*heartbeat.SendHeartbeatResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(SEND_HEARTBEAT_URL, req))

	if resp == nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		// try to parse the returning JSON into the response
		var sendHeartbeatResp heartbeat.SendHeartbeatResponse
		if err = resp.Body.FromJsonTo(&sendHeartbeatResp); err != nil {
			message := "Server response can not be parsed, " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		// parsed successfuly with no errors
		return &sendHeartbeatResp, nil
	}
}
