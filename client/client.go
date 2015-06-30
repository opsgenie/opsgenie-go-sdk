// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

/*
	Package client manages the creation of API clients.
	API user first creates a pointer of type OpsGenieClient. Following that
	he/she can set some configurations for HTTP communication layer by setting
	a proxy definition and/or transport layer options.

	Introduction

	The most fundamental and general use case is being able to access the
	OpsGenie Web API by coding a Go program.
	The program -by mean of a client application- can send OpsGenie Web API
	the requests using the 'client' package in a higher level. For the programmer
	of the client application, that reduces the number of LoCs.
	Besides it will result a less error-prone application and reduce
	the complexity by hiding the low-level networking, error-handling and
	byte-processing calls.

	Package client has ports for all entry points to the Web API.
	The OpsGenie Web API is structured in JSON-bodied
	calls (except the file attachment).
*/
package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/franela/goreq"
	goquery "github.com/google/go-querystring/query"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
)

// OpsGenie Go SDK performs HTTP calls to the Web API.
// The Web API is designated by a URL so called an endpoint
var ENDPOINT_URL = "https://api.opsgenie.com"

const DEFAULT_CONNECTION_TIMEOUT_IN_SECONDS time.Duration = 50 * time.Second
const DEFAULT_REQUEST_TIMEOUT_IN_SECONDS time.Duration = 100 * time.Second
const DEFAULT_MAX_RETRY_ATTEMPTS int = 5
const TIME_SLEEP_BETWEEN_REQUESTS time.Duration = 500 * time.Millisecond

// User-Agent values tool/version (OS;GO_Version;language)
type RequestHeaderUserAgent struct {
	sdkName   string
	version   string
	os        string
	goVersion string
	timezone  string
}

func (p RequestHeaderUserAgent) ToString() string {
	return fmt.Sprintf("%s/%s (%s;%s;%s)", p.sdkName, p.version, p.os, p.goVersion, p.timezone)
}

var userAgentParam RequestHeaderUserAgent

// OpsGenieClient is a general data type used for:
// 	- authenticating callers through their api keys and
// 	- instantiating "alert" and "heartbeat" clients
//	- setting HTTP transport layer configurations
type OpsGenieClient struct {
	proxy                 *ClientProxyConfiguration
	httpTransportSettings *HttpTransportSettings
	apiKey                string
	opsGenieApiUrl        string
}

// Setters
func (cli *OpsGenieClient) SetClientProxyConfiguration(conf *ClientProxyConfiguration) {
	cli.proxy = conf
}

func (cli *OpsGenieClient) SetHttpTransportSettings(settings *HttpTransportSettings) {
	cli.httpTransportSettings = settings
}

func (cli *OpsGenieClient) SetApiKey(key string) {
	cli.apiKey = key
}

func (cli *OpsGenieClient) SetOpsGenieApiUrl(url string) {
	if url != "" {
		cli.opsGenieApiUrl = url
	}
}

func (cli *OpsGenieClient) GetOpsGenieApiUrl() string {
	if cli.opsGenieApiUrl == "" {
		cli.opsGenieApiUrl = ENDPOINT_URL
	}
	return cli.opsGenieApiUrl
}

func (cli *OpsGenieClient) GetApiKey() string {
	return cli.apiKey
}

func (cli *OpsGenieClient) MakeHttpTransportSettings() {
	if cli.httpTransportSettings != nil {
		if cli.httpTransportSettings.MaxRetryAttempts <= 0 {
			cli.httpTransportSettings.MaxRetryAttempts = DEFAULT_MAX_RETRY_ATTEMPTS
		}
		if cli.httpTransportSettings.ConnectionTimeout <= 0 {
			cli.httpTransportSettings.ConnectionTimeout = DEFAULT_CONNECTION_TIMEOUT_IN_SECONDS
		}
		if cli.httpTransportSettings.RequestTimeout <= 0 {
			cli.httpTransportSettings.RequestTimeout = DEFAULT_REQUEST_TIMEOUT_IN_SECONDS
		}
	} else {
		cli.httpTransportSettings = &HttpTransportSettings{MaxRetryAttempts: DEFAULT_MAX_RETRY_ATTEMPTS, ConnectionTimeout: DEFAULT_CONNECTION_TIMEOUT_IN_SECONDS, RequestTimeout: DEFAULT_REQUEST_TIMEOUT_IN_SECONDS}
	}
}

// Instantiates a new OpsGenieAlertClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Alert() (*OpsGenieAlertClient, error) {
	cli.MakeHttpTransportSettings()

	alertClient := new(OpsGenieAlertClient)
	alertClient.SetOpsGenieClient(*cli)

	if cli.opsGenieApiUrl == "" {
		alertClient.SetOpsGenieApiUrl(ENDPOINT_URL)
	}

	return alertClient, nil
}

// Instantiates a new OpsGenieHeartbeatClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Heartbeat() (*OpsGenieHeartbeatClient, error) {
	cli.MakeHttpTransportSettings()

	heartbeatClient := new(OpsGenieHeartbeatClient)
	heartbeatClient.SetOpsGenieClient(*cli)

	if cli.opsGenieApiUrl == "" {
		heartbeatClient.SetOpsGenieApiUrl(ENDPOINT_URL)
	}

	return heartbeatClient, nil
}

// Instantiates a new OpsGenieIntegrationClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Integration() (*OpsGenieIntegrationClient, error) {
	cli.MakeHttpTransportSettings()

	integrationClient := new(OpsGenieIntegrationClient)
	integrationClient.SetOpsGenieClient(*cli)

	if cli.opsGenieApiUrl == "" {
		integrationClient.SetOpsGenieApiUrl(ENDPOINT_URL)
	}

	return integrationClient, nil
}

// Instantiates a new OpsGeniePolicyClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Policy() (*OpsGeniePolicyClient, error) {
	cli.MakeHttpTransportSettings()

	policyClient := new(OpsGeniePolicyClient)
	policyClient.SetOpsGenieClient(*cli)

	if cli.opsGenieApiUrl == "" {
		policyClient.SetOpsGenieApiUrl(ENDPOINT_URL)
	}

	return policyClient, nil
}

func (cli *OpsGenieClient) buildCommonRequestProps() goreq.Request {
	if cli.httpTransportSettings == nil {
		cli.MakeHttpTransportSettings()
	}
	goreq.SetConnectTimeout(cli.httpTransportSettings.ConnectionTimeout)
	req := goreq.Request{}
	if cli.proxy != nil {
		req.Proxy = cli.proxy.ToString()
	}
	req.UserAgent = userAgentParam.ToString()
	req.Timeout = cli.httpTransportSettings.RequestTimeout
	req.Insecure = true
	//	req.AddHeader("Connection","Keep-Alive")

	return req
}

func (cli *OpsGenieClient) buildGetRequest(uri string, request interface{}) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "GET"
	req.ContentType = "application/x-www-form-urlencoded; charset=UTF-8"
	uri = cli.opsGenieApiUrl + uri
	v, _ := goquery.Values(request)
	req.Uri = uri + "?" + v.Encode()
	logging.Logger().Info("Executing OpsGenie request to ["+uri+"] with parameters: ", v)
	return req
}

func (cli *OpsGenieClient) buildPostRequest(uri string, request interface{}) goreq.Request {
	req := cli.buildCommonRequestProps()
	req.Method = "POST"
	req.ContentType = "application/json; charset=utf-8"
	req.Uri = cli.opsGenieApiUrl + uri
	req.Body = request
	j, _ := json.Marshal(request)
	logging.Logger().Info("Executing OpsGenie request to ["+req.Uri+"] with content parameters: ", string(j))

	return req
}

func (cli *OpsGenieClient) buildDeleteRequest(uri string, request interface{}) goreq.Request {
	req := cli.buildGetRequest(uri, request)
	req.Method = "DELETE"
	return req
}

func (cli *OpsGenieClient) sendRequest(req goreq.Request) (*goreq.Response, error) {
	// send the request
	var resp *goreq.Response
	var err error
	for i := 0; i < cli.httpTransportSettings.MaxRetryAttempts; i++ {
		resp, err = req.Do()
		if err == nil && resp.StatusCode < 500 {
			break
		}
		if resp != nil {
			defer resp.Body.Close()
			logging.Logger().Info(fmt.Sprintf("Retrying request [%s] ResponseCode:[%d]. RetryCount: %d", req.Uri, resp.StatusCode, (i + 1)))
		} else {
			logging.Logger().Info(fmt.Sprintf("Retrying request [%s] Reason:[%s]. RetryCount: %d", req.Uri, err.Error(), (i + 1)))
		}
		time.Sleep(TIME_SLEEP_BETWEEN_REQUESTS * time.Duration(i+1))
	}
	if err != nil {
		message := "Unable to send the request " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	// check for the returning http status
	statusCode := resp.StatusCode
	if statusCode >= 400 {
		body, err := resp.Body.ToString()
		if err != nil {
			message := "Server response with error can not be parsed " + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return nil, getErrorMessage(statusCode, body)
	} else {
		return resp, nil
	}
}

func getErrorMessage(httpStatusCode int, responseBody string) error {
	if httpStatusCode >= 400 && httpStatusCode < 500 {
		message := fmt.Sprintf("Client error occurred; Response Code: %d, Response Body: "+responseBody, httpStatusCode)
		logging.Logger().Warn(message)
		return errors.New(message)
	}
	if httpStatusCode >= 500 {
		message := fmt.Sprintf("Server error occurred; Response Code: %d, Response Body: "+responseBody, httpStatusCode)
		logging.Logger().Info(message)
		return errors.New(message)
	}
	return nil
}

// Initializer for the package client
// Initializes the User-Agent parameter of the requests.
// TODO version information must be read from a MANIFEST file
func init() {
	userAgentParam.sdkName = "opsgenie-go-sdk"
	userAgentParam.version = "1.0.0"
	userAgentParam.os = runtime.GOOS
	userAgentParam.goVersion = runtime.Version()
	userAgentParam.timezone = time.Local.String()
}
