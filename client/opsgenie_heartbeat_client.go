package client

import (
	"errors"

	"github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
	goquery "github.com/google/go-querystring/query"
	"github.com/franela/goreq"
	"encoding/json"
)

const (
	addHeartbeatURL     = "/v1/json/heartbeat"
	updateHeartbeatURL  = "/v1/json/heartbeat"
	enableHeartbeatURL  = "/v1/json/heartbeat/enable"
	disableHeartbeatURL = "/v1/json/heartbeat/disable"
	deleteHeartbeatURL  = "/v1/json/heartbeat"
	getHeartbeatURL     = "/v1/json/heartbeat"
	listHeartbeatURL    = "/v1/json/heartbeat"
	sendHeartbeatURL    = "/v1/json/heartbeat/send"

	heartbeatUrlV2 = "/v2/heartbeats"
	heartbeatEnablePath = "/enable"
	heartbeatDisablePath = "/disable"
	heartbeatPingPath = "/ping"

)

// OpsGenieHeartbeatClient is the data type to make Heartbeat API requests.
type OpsGenieHeartbeatClient struct {
	OpsGenieClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieHeartbeatClient.
func (cli *OpsGenieHeartbeatClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// Add method creates a heartbeat at OpsGenie.
// OLD METHOD (ADD)
/*func (cli *OpsGenieHeartbeatClient) Add(req heartbeat.AddHeartbeatRequest) (*heartbeat.AddHeartbeatResponse, error) {
	req.APIKey = cli.apiKey

	resp, err := cli.sendRequest(cli.buildPostRequest(addHeartbeatURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var addHeartbeatResp heartbeat.AddHeartbeatResponse
	if err = resp.Body.FromJsonTo(&addHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &addHeartbeatResp, nil
}*/

// Add method creates a heartbeat at OpsGenie.
// NEW METHOD (ADD)
func (cli *OpsGenieHeartbeatClient) Add(req heartbeat.AddHeartbeatRequest) (*heartbeat.AddHeartbeatResponse, error) {
	resp, err := cli.sendRequest(cli.buildPostRequestForHb(heartbeatUrlV2, convertAddToV2Request(req)))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var addHeartbeatResp heartbeat.HeartbeatResponseV2
	if err = resp.Body.FromJsonTo(&addHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}


	result := convertAddToV1Response(&addHeartbeatResp)
	result.Code = resp.StatusCode

	return result, nil
}

// Update method changes configuration of an existing heartbeat at OpsGenie.
// OLD METHOD (UPDATE)
/*func (cli *OpsGenieHeartbeatClient) Update(req heartbeat.UpdateHeartbeatRequest) (*heartbeat.UpdateHeartbeatResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(updateHeartbeatURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updateHeartbeatResp heartbeat.UpdateHeartbeatResponse
	if err = resp.Body.FromJsonTo(&updateHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &updateHeartbeatResp, nil
}*/

// Update method changes configuration of an existing heartbeat at OpsGenie.
// NEW METHOD (UPDATE)
func (cli *OpsGenieHeartbeatClient) Update(req heartbeat.UpdateHeartbeatRequest) (*heartbeat.UpdateHeartbeatResponse, error) {
	newReq := convertUpdateToV2Request(req)
	name := req.Name
	if len(name) == 0 {
		panic("Name can not be empty.")
	}
	resp, err := cli.sendRequest(cli.buildUpdateRequestForHb(heartbeatUrlV2, name, newReq))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updateHeartbeatResp heartbeat.HeartbeatMetaResponseV2
	if err = resp.Body.FromJsonTo(&updateHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	result := convertUpdateToV1Response(&updateHeartbeatResp)
	result.Code = resp.StatusCode

	return result, nil
}

// Enable method enables an heartbeat at OpsGenie.
func (cli *OpsGenieHeartbeatClient) Enable(req heartbeat.EnableHeartbeatRequest) (*heartbeat.EnableHeartbeatResponse, error) {
	resp, err := cli.sendRequest(cli.buildPostRequestForHbAction(heartbeatUrlV2, heartbeatEnablePath, req.Name))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var enableHeartbeatResp heartbeat.HeartbeatMetaResponseV2
	if err = resp.Body.FromJsonTo(&enableHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	result := heartbeat.EnableHeartbeatResponse{}
	result.Code = resp.StatusCode

	return &result, nil
}

// Disable method disables an heartbeat at OpsGenie.
func (cli *OpsGenieHeartbeatClient) Disable(req heartbeat.DisableHeartbeatRequest) (*heartbeat.DisableHeartbeatResponse, error) {
	resp, err := cli.sendRequest(cli.buildPostRequestForHbAction(heartbeatUrlV2, heartbeatDisablePath, req.Name))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var disableHeartbeatResp heartbeat.HeartbeatMetaResponseV2
	if err = resp.Body.FromJsonTo(&disableHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	result := heartbeat.DisableHeartbeatResponse{}
	result.Code = resp.StatusCode

	return &result, nil

}

// Delete method deletes an heartbeat from OpsGenie.
func (cli *OpsGenieHeartbeatClient) Delete(req heartbeat.DeleteHeartbeatRequest) (*heartbeat.DeleteHeartbeatResponse, error) {
	hbName := req.Name
	resp, err := cli.sendRequest(cli.buildDeleteRequestForHb(heartbeatUrlV2, hbName))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deleteHeartbeatResp heartbeat.DeleteHeartbeatResponse
	if err = resp.Body.FromJsonTo(&deleteHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	result := heartbeat.DeleteHeartbeatResponse{}
	result.Code = resp.StatusCode

	return &result, nil
}

// Get method retrieves an heartbeat with details from OpsGenie.
// OLD METHOD (GET)
/*func (cli *OpsGenieHeartbeatClient) Get(req heartbeat.GetHeartbeatRequest) (*heartbeat.GetHeartbeatResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(getHeartbeatURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getHeartbeatResp heartbeat.GetHeartbeatResponse
	if err = resp.Body.FromJsonTo(&getHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &getHeartbeatResp, nil
}*/

// Get method retrieves an heartbeat with details from OpsGenie.
// NEW METHOD (GET)
func (cli *OpsGenieHeartbeatClient) Get(req heartbeat.GetHeartbeatRequest) (*heartbeat.GetHeartbeatResponse, error) {
	resp, err := cli.sendRequest(cli.buildGetRequestForHb(heartbeatUrlV2, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getHeartbeatResp heartbeat.HeartbeatResponseV2
	if err = resp.Body.FromJsonTo(&getHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return convertGetToV1Response(&getHeartbeatResp), nil
}

// List method retrieves heartbeats from OpsGenie.
func (cli *OpsGenieHeartbeatClient) List(req heartbeat.ListHeartbeatsRequest) (*heartbeat.ListHeartbeatsResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(listHeartbeatURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var listHeartbeatsResp heartbeat.ListHeartbeatsResponse
	if err = resp.Body.FromJsonTo(&listHeartbeatsResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &listHeartbeatsResp, nil
}

// Send method sends an Heartbeat Signal to OpsGenie.
func (cli *OpsGenieHeartbeatClient) Send(req heartbeat.SendHeartbeatRequest) (*heartbeat.SendHeartbeatResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(sendHeartbeatURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sendHeartbeatResp heartbeat.SendHeartbeatResponse
	if err = resp.Body.FromJsonTo(&sendHeartbeatResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &sendHeartbeatResp, nil
}

func (cli *OpsGenieHeartbeatClient) Ping(req heartbeat.PingHeartbeatRequest) (*heartbeat.PingHeartbeatResponse, error) {
	hbName := req.Name
	resp, err := cli.sendRequest(cli.buildPingRequestForHb(heartbeatUrlV2, hbName))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pingHeartbeatResponse heartbeat.PingHeartbeatResponse
	if err = resp.Body.FromJsonTo(&pingHeartbeatResponse); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &pingHeartbeatResponse, nil
}

func (cli *OpsGenieClient) buildPostRequestForHb(uri string, request interface{}) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "POST"
	req.ContentType = "application/json; charset=utf-8"
	req.AddHeader("Authorization", authHeader + cli.apiKey)
	req.Uri = cli.OpsGenieAPIUrl() + uri
	req.Body = request
	j, _ := json.Marshal(request)
	logging.Logger().Info("Executing OpsGenie request to ["+req.Uri+"] with content parameters: ", string(j))

	return req
}

func convertAddToV2Request(requestV1 heartbeat.AddHeartbeatRequest) heartbeat.AddHeartbeatRequestV2  {
	hbName := requestV1.Name
	hbInterval := requestV1.Interval
	hbIntervalUnit := requestV1.IntervalUnit
	if len(hbIntervalUnit) == 0 {
		hbIntervalUnit = "minutes"
	}
	hbEnabled := requestV1.Enabled
	hbDescription := requestV1.Description

	var result = heartbeat.AddHeartbeatRequestV2{}

	result.Name = hbName
	result.Interval = hbInterval
	result.IntervalUnit = hbIntervalUnit
	result.Enabled = *hbEnabled
	if hbEnabled == nil {
		result.Enabled = true
	}
	result.Description = hbDescription

	return result
}

func convertAddToV1Response(responseV2 *heartbeat.HeartbeatResponseV2) *heartbeat.AddHeartbeatResponse {
	data := responseV2.Data

	var result = heartbeat.AddHeartbeatResponse{}

	result.Name = data.Name

	return &result
}

func (cli *OpsGenieClient) buildUpdateRequestForHb(uri string, name string, request interface{}) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "PATCH"
	req.ContentType = "application/json; charset=utf-8"
	req.AddHeader("Authorization", authHeader + cli.apiKey)
	req.Uri = cli.OpsGenieAPIUrl() + uri + "/" + name
	req.Body = request
	j, _ := json.Marshal(request)
	logging.Logger().Info("Executing OpsGenie request to ["+req.Uri+"] with content parameters: ", string(j))

	return req
}

func convertUpdateToV2Request(requestV1 heartbeat.UpdateHeartbeatRequest) heartbeat.UpdateHeartbeatRequestV2  {
	hbInterval := requestV1.Interval
	hbIntervalUnit := requestV1.IntervalUnit
	hbEnabled := requestV1.Enabled
	hbDescription := requestV1.Description

	var result = heartbeat.UpdateHeartbeatRequestV2{}

	result.Interval = hbInterval
	if len(hbIntervalUnit) != 0 {
		result.IntervalUnit = hbIntervalUnit
	}
	if len(hbDescription) != 0 {
		result.Description = hbDescription
	}
	if hbEnabled != nil {
		result.Enabled = *hbEnabled
	}

	return result
}

func convertUpdateToV1Response(responseV2 *heartbeat.HeartbeatMetaResponseV2) *heartbeat.UpdateHeartbeatResponse {
	data := responseV2.Data

	var result = heartbeat.UpdateHeartbeatResponse{}

	result.Name = data.Name

	return &result
}

func (cli *OpsGenieHeartbeatClient) buildGetRequestForHb(uri string, request interface{}) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "GET"
	req.ContentType = "application/x-www-form-urlencoded; charset=UTF-8"
	req.AddHeader("Authorization", authHeader + cli.apiKey)
	uri = cli.OpsGenieAPIUrl() + uri
	if request != nil {
		v, _ := goquery.Values(request)
		req.Uri = uri + "/" + v.Get("name")
	} else {
		req.Uri = uri
	}
	logging.Logger().Info("Executing OpsGenie request to [" + uri + "] with parameters: ")
	return req
}

func convertGetToV1Response(responseV2 *heartbeat.HeartbeatResponseV2) *heartbeat.GetHeartbeatResponse {
	data := responseV2.Data

	var result = heartbeat.GetHeartbeatResponse{}

	result.Name = data.Name
	result.Description = data.Description
	result.Interval = data.Interval
	result.IntervalUnit = data.IntervalUnit
	result.Enabled = data.Enabled
	result.Expired = data.Expired

	return &result
}

func (cli *OpsGenieClient) buildPostRequestForHbAction(uri string, actionUri string, hbName string) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "POST"
	req.ContentType = "application/json; charset=utf-8"
	req.AddHeader("Authorization", authHeader + cli.apiKey)
	if len(hbName) == 0 {
		panic("[name] field can not be empty.")
	}
	req.Uri = cli.OpsGenieAPIUrl() + uri + "/" + hbName + actionUri
	logging.Logger().Info("Executing OpsGenie request to ["+req.Uri+"] .")

	return req
}

func (cli *OpsGenieClient) buildDeleteRequestForHb(uri string, hbName string) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "DELETE"
	req.ContentType = "application/json; charset=utf-8"
	req.AddHeader("Authorization", authHeader + cli.apiKey)
	if len(hbName) == 0 {
		panic("[name] field can not be empty.")
	}
	req.Uri = cli.OpsGenieAPIUrl() + uri + "/" + hbName
	logging.Logger().Info("Executing OpsGenie request to ["+req.Uri+"] .")

	return req
}

func (cli *OpsGenieClient) buildPingRequestForHb(uri string, hbName string) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "POST"
	req.ContentType = "application/json; charset=utf-8"
	req.AddHeader("Authorization", authHeader + cli.apiKey)
	if len(hbName) == 0 {
		panic("[name] field can not be empty.")
	}
	req.Uri = cli.OpsGenieAPIUrl() + uri + "/" + hbName + heartbeatPingPath
	logging.Logger().Info("Executing OpsGenie request to ["+req.Uri+"] .")

	return req
}

