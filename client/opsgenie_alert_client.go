// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by an Apache Software 
// license that can be found in the LICENSE file.
package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/alerts"
	"io"
	"os"
	"fmt"
	"mime/multipart"
	"net/http"
	"bytes"
	"errors"
	"net/url"
	"time"
	"io/ioutil"
	"crypto/tls"
	"net"
	"path/filepath"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
)

const(
	CREATE_ALERT_URL 			= "/v1/json/alert"
	CLOSE_ALERT_URL 			= "/v1/json/alert/close"
	DELETE_ALERT_URL 			= "/v1/json/alert"
	GET_ALERT_URL 				= "/v1/json/alert"
	LIST_ALERTS_URL 			= "/v1/json/alert"
	LIST_ALERT_NOTES_URL 		= "/v1/json/alert/note"
	LIST_ALERT_LOGS_URL 		= "/v1/json/alert/log"
	LIST_ALERT_RECIPIENTS_URL 	= "/v1/json/alert/recipient"
	ACKNOWLEDGE_ALERT_URL 		= "/v1/json/alert/acknowledge"
	RENOTIFY_ALERT_URL 			= "/v1/json/alert/renotify"
	TAKE_OWNERSHIP_ALERT_URL 	= "/v1/json/alert/takeOwnership"
	ASSIGN_OWNERSHIP_ALERT_URL 	= "/v1/json/alert/assign"
	ADD_TEAM_ALERT_URL 			= "/v1/json/alert/team"
	ADD_RECIPIENT_ALERT_URL 	= "/v1/json/alert/recipient"
	ADD_NOTE_ALERT_URL 			= "/v1/json/alert/note"
	EXECUTE_ACTION_ALERT_URL 	= "/v1/json/alert/executeAction"
	ATTACH_FILE_ALERT_URL 		= "/v1/json/alert/attach"
)

type OpsGenieAlertClient struct {
	OpsGenieClient
}

func (cli *OpsGenieAlertClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieAlertClient) Create(req alerts.CreateAlertRequest) (*alerts.CreateAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(CREATE_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var createAlertResp alerts.CreateAlertResponse
		// check if the response can be unmarshalled
		if err = resp.Body.FromJsonTo(&createAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &createAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) Close(req alerts.CloseAlertRequest) (*alerts.CloseAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(CLOSE_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var closeAlertResp alerts.CloseAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&closeAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &closeAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) Delete(req alerts.DeleteAlertRequest) (*alerts.DeleteAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildDeleteRequest(DELETE_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var deleteAlertResp alerts.DeleteAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&deleteAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &deleteAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) Get(req alerts.GetAlertRequest) (*alerts.GetAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(GET_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var getAlertResp alerts.GetAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&getAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &getAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) List(req alerts.ListAlertsRequest) (*alerts.ListAlertsResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(LIST_ALERTS_URL, req))

	if resp == nil {
		return nil, errors.New(err.Error())
	}else{
		defer resp.Body.Close()

		var listAlertsResp alerts.ListAlertsResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&listAlertsResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &listAlertsResp, nil
	}
}

func (cli *OpsGenieAlertClient) ListNotes(req alerts.ListAlertNotesRequest) (*alerts.ListAlertNotesResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(LIST_ALERT_NOTES_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var listAlertNotesResp alerts.ListAlertNotesResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&listAlertNotesResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &listAlertNotesResp, nil
	}
}

func (cli *OpsGenieAlertClient) ListLogs(req alerts.ListAlertLogsRequest) (*alerts.ListAlertLogsResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(LIST_ALERT_LOGS_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var listAlertLogsResp alerts.ListAlertLogsResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&listAlertLogsResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &listAlertLogsResp, nil
	}
}

func (cli *OpsGenieAlertClient) ListRecipients(req alerts.ListAlertRecipientsRequest) (*alerts.ListAlertRecipientsResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(LIST_ALERT_RECIPIENTS_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var listAlertRecipientsResp alerts.ListAlertRecipientsResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&listAlertRecipientsResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &listAlertRecipientsResp, nil
	}
}

func (cli *OpsGenieAlertClient) Acknowledge(req alerts.AcknowledgeAlertRequest) (*alerts.AcknowledgeAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ACKNOWLEDGE_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var acknowledgeAlertResp alerts.AcknowledgeAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&acknowledgeAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &acknowledgeAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) Renotify(req alerts.RenotifyAlertRequest) (*alerts.RenotifyAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(RENOTIFY_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var renotifyAlertResp alerts.RenotifyAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&renotifyAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &renotifyAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) TakeOwnership(req alerts.TakeOwnershipAlertRequest) (*alerts.TakeOwnershipAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(TAKE_OWNERSHIP_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var takeOwnershipResp alerts.TakeOwnershipAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&takeOwnershipResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &takeOwnershipResp, nil
	}
}

func (cli *OpsGenieAlertClient) AssignOwner(req alerts.AssignOwnerAlertRequest) (*alerts.AssignOwnerAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ASSIGN_OWNERSHIP_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var assignOwnerAlertResp alerts.AssignOwnerAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&assignOwnerAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &assignOwnerAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) AddTeam(req alerts.AddTeamAlertRequest) (*alerts.AddTeamAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ADD_TEAM_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var addTeamAlertResp alerts.AddTeamAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&addTeamAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &addTeamAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) AddRecipient(req alerts.AddRecipientAlertRequest) (*alerts.AddRecipientAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ADD_RECIPIENT_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var addRecipientAlertResp alerts.AddRecipientAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&addRecipientAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &addRecipientAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) AddNote(req alerts.AddNoteAlertRequest) (*alerts.AddNoteAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(ADD_NOTE_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var addNoteAlertResp alerts.AddNoteAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&addNoteAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &addNoteAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) ExecuteAction(req alerts.ExecuteActionAlertRequest) (*alerts.ExecuteActionAlertResponse, error) {
	req.ApiKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(EXECUTE_ACTION_ALERT_URL, req))

	if resp == nil {
		return nil, err
	}else{
		defer resp.Body.Close()

		var executeActionAlertResp alerts.ExecuteActionAlertResponse
		// check if the response can be unmarshalled
		if err =  resp.Body.FromJsonTo(&executeActionAlertResp); err != nil {
			message := "Server response can not be parsed, "  + err.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		return &executeActionAlertResp, nil
	}
}

func (cli *OpsGenieAlertClient) AttachFile(req alerts.AttachFileAlertRequest) (*alerts.AttachFileAlertResponse, error) {
	req.ApiKey = cli.apiKey
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	path := req.Attachment.Name()
	file, err := os.Open(path)
	defer file.Close()
	if err != nil{
		message := "Attachment can not be opened for reading. " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	// add the attachment
	fw, err := w.CreateFormFile("attachment", filepath.Base(path))
	if err != nil {
		message := "Can not build the request with the field attachment. " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	if _, err := io.Copy(fw, file); err != nil {
		message := "Can not copy the attachment into the request. " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	// Add the other fields
	// empty fields should not be placed into the request
	// otherwise it yields an incomplete boundary exception
	if req.ApiKey != ""{
		if err = writeField(*w, "apiKey", req.ApiKey); err != nil{
			return nil, err
		}
	}
	if req.Id != ""{
		if err = writeField(*w, "id", req.Id); err != nil {
			return nil, err
		}
	}
	if req.Alias != ""{
		if err = writeField(*w, "alias", req.Alias); err != nil{
			return nil, err
		}
	}
	if req.User != ""{
		if err = writeField(*w, "user", req.User); err != nil {
			return nil, err
		}
	}
	if req.Source != "" {
		if err = writeField(*w, "source", req.Source); err != nil {
			return nil, err
		}
	}
	if req.IndexFile != "" {
		if err = writeField(*w, "indexFile", req.IndexFile); err != nil {
			return nil, err
		}
	}
	if req.Note != ""{
		if err = writeField(*w, "note", req.Note); err != nil {
			return nil, err
		}
	}

	w.Close()
	httpReq, err := http.NewRequest("POST", cli.opsGenieApiUrl + ATTACH_FILE_ALERT_URL, &b)
	if err != nil {
		message := "Can not create the multipart/form-data request. " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	httpReq.Header.Set("Content-Type", w.FormDataContentType())
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
		Proxy: http.ProxyFromEnvironment,
		Dial:  func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, cli.httpTransportSettings.ConnectionTimeout)
			if err != nil {
				message := "Error occurred while connecting: " + err.Error()
				logging.Logger().Warn(message)
				return nil, errors.New(message)
			}
			conn.SetDeadline(time.Now().Add(cli.httpTransportSettings.RequestTimeout))
			return conn, nil
		},
	}
	client := &http.Client{Transport: transport}
	// proxy settings
	if cli.proxy != nil {
		proxyUrl, proxyErr := url.Parse(cli.proxy.ToString())
		if proxyErr != nil {
			message := "Can not set the proxy configuration "  + proxyErr.Error()
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	url := httpReq.URL.String()
	logging.Logger().Info("Executing OpsGenie request to [" + url + "] with multipart data.")
	var res *http.Response
	for i := 0; i < cli.httpTransportSettings.MaxRetryAttempts; i++ {
		res, err = client.Do(httpReq)
		if err == nil {
			defer res.Body.Close()
			break
		}
		if res != nil{
			logging.Logger().Info(fmt.Sprintf("Retrying request [%s] ResponseCode:[%d]. RetryCount: %d",url, res.StatusCode,(i+1)))
		}else{
			logging.Logger().Info(fmt.Sprintf("Retrying request [%s] Reason:[%s]. RetryCount: %d",url, err.Error(),(i+1)))
		}
		time.Sleep(TIME_SLEEP_BETWEEN_REQUESTS)
	}

	if err != nil {
		message := "Can not attach the file, unable to send the request. " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	// check the returning HTTP status code
	httpStatusCode := res.StatusCode
	if httpStatusCode >= 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err == nil{
			return nil, getErrorMessage(httpStatusCode, string(body[:]))
		}else{
			message := fmt.Sprintf("Couldn't read the response, ", err)
			logging.Logger().Warn(message)
			return nil, errors.New(message)
		}
	}

	attachFileAlertResp := alerts.AttachFileAlertResponse{Status: res.Status, Code: res.StatusCode}
	return &attachFileAlertResp, nil
}

func writeField (w multipart.Writer, fieldName string, fieldVal string) error {
	if err := w.WriteField(fieldName, fieldVal); err != nil {
		message := "Can not write field "  + fieldName +" into the request. Reason: " + err.Error()
		logging.Logger().Warn(message)
		return errors.New(message)
	}
	return nil
}