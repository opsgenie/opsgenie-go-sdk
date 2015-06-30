package test

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
	"time"
)

type AlertTestSuite struct {
	suite.Suite
}

func (suite *AlertTestSuite) SetupTest() {
	deleteAllAlerts(suite.T())
}

func (suite *AlertTestSuite) TestCreateAlert() {
	t := suite.T()
	suffix := time.Now().String()
	id := createAlert(t, suffix)
	//get with id
	getAlertReq := alerts.GetAlertRequest{Id: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, make(map[string]interface{}))

	t.Log("[OK] get alert with id successful")
	tinyId := getResp.TinyId

	//get with tinyid
	getAlertReq = alerts.GetAlertRequest{TinyId: tinyId}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, make(map[string]interface{}))

	t.Log("[OK] get alert with tinyid successful")

	alias := getResp.Alias

	//get with alias
	getAlertReq = alerts.GetAlertRequest{Alias: alias}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, make(map[string]interface{}))

	t.Log("[OK] get alert with alias successful")
}

func (suite *AlertTestSuite) TestDeleteAlert() {
	t := suite.T()
	suffix := time.Now().String()
	id := createAlert(t, suffix)
	delreq := alerts.DeleteAlertRequest{Id: id, Source: "Go API Test"}
	delResp, alertErr := cli.Delete(delreq)

	require.Nil(t, alertErr)
	require.NotNil(t, delResp)
	require.Equal(t, 200, delResp.Code, "Response Code should be 200")

	getAlertReq := alerts.GetAlertRequest{Id: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.NotNil(t, alertErr)
	require.Nil(t, getResp)
	require.True(t, strings.Contains(alertErr.Error(), "Alert with id ["+id+"] does not exist"), "Error should contains ")
	t.Log("[OK] alert deleted")

	//test delete alert with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	delreq = alerts.DeleteAlertRequest{Alias: "alert" + suffix, Source: "Go API Test"}
	delResp, alertErr = cli.Delete(delreq)

	require.Nil(t, alertErr)
	require.NotNil(t, delResp)
	require.Equal(t, 200, delResp.Code, "Response Code should be 200")

	getAlertReq = alerts.GetAlertRequest{Alias: "alert" + suffix}
	getResp, alertErr = cli.Get(getAlertReq)

	require.NotNil(t, alertErr)
	require.Nil(t, getResp)
	require.True(t, strings.Contains(alertErr.Error(), "Alert with alias [alert"+suffix+"] does not exist"), "Error should contains ")
	t.Log("[OK] alert deleted")
}

func (suite *AlertTestSuite) TestCloseAlert() {
	t := suite.T()
	suffix := time.Now().String()
	id := createAlert(t, suffix)
	closeReq := alerts.CloseAlertRequest{Id: id, Source: "Go API Test", Note: "closing the alert"}
	closeResp, alertErr := cli.Close(closeReq)

	require.Nil(t, alertErr)
	require.NotNil(t, closeResp)
	require.Equal(t, 200, closeResp.Code, "Response Code should be 200")

	getAlertReq := alerts.GetAlertRequest{Id: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"status": "closed", "isSeen": true})
	t.Log("[OK] alert closed")

	//	//test close alert with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	closeReq = alerts.CloseAlertRequest{Alias: "alert" + suffix, Source: "Go API Test"}
	closeResp, alertErr = cli.Close(closeReq)

	require.Nil(t, alertErr)
	require.NotNil(t, closeResp)
	require.Equal(t, 200, closeResp.Code, "Response Code should be 200")

	getAlertReq = alerts.GetAlertRequest{Id: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"status": "closed", "isSeen": true})
	t.Log("[OK] alert closed")
}

func (suite *AlertTestSuite) TestAcknowledgeAlert() {
	t := suite.T()
	suffix := time.Now().String()
	id := createAlert(t, suffix)
	owner := "owneruser"
	ackReq := alerts.AcknowledgeAlertRequest{Id: id, User: owner}
	ackResponse, alertErr := cli.Acknowledge(ackReq)

	require.Nil(t, alertErr)
	require.NotNil(t, ackResponse)
	require.Equal(t, 200, ackResponse.Code, "Response Code should be 200")

	getAlertReq := alerts.GetAlertRequest{Id: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"acknowledged": true, "isSeen": true, "owner": owner})
	t.Log("[OK] alert acked")

	//test close alert with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	ackReq = alerts.AcknowledgeAlertRequest{Alias: "alert" + suffix, User: owner, Source: "Go API Test"}
	ackResponse, alertErr = cli.Acknowledge(ackReq)

	require.Nil(t, alertErr)
	require.NotNil(t, ackResponse)
	require.Equal(t, 200, ackResponse.Code, "Response Code should be 200")

	getAlertReq = alerts.GetAlertRequest{Id: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"acknowledged": true, "isSeen": true, "owner": owner})
	t.Log("[OK] alert acked")
}

func assertAlert(t *testing.T, id string, response *alerts.GetAlertResponse, suffix string, overwrite map[string]interface{}) {
	require.Equal(t, id, response.Id, "Alert Id comparison failed")
	require.Equal(t, "message"+suffix, response.Message, "Message comparison failed")
	require.Equal(t, "description"+suffix, response.Description, "Description comparison failed")
	require.Equal(t, "source"+suffix, response.Source, "Source comparison failed")
	if val, contains := overwrite["owner"]; contains {
		require.Equal(t, val.(string), response.Owner, "Owner comparison failed")
	} else {
		require.Equal(t, "", response.Owner, "Owner comparison failed")
	}
	require.Equal(t, "alert"+suffix, response.Alias, "Alias comparison failed")
	require.Equal(t, "entity"+suffix, response.Entity, "Entity comparison failed")
	if val, contains := overwrite["status"]; contains {
		require.Equal(t, val.(string), response.Status, "Status comparison failed")
	} else {
		require.Equal(t, "open", response.Status, "Status comparison failed")
	}
	require.Equal(t, testCfg.Alert.Actions, response.Actions, "Action comparison failed")
	require.Equal(t, []string{"tag1", "tag2", "tag3"}, response.Tags, "Tags comparison failed")
	if val, contains := overwrite["recipients"]; contains {
		require.Equal(t, val, response.Recipients, "recipients comparison failed")
	} else {
		require.Equal(t, []string{"user1"}, response.Recipients, "recipients comparison failed")
	}

	if val, contains := overwrite["teams"]; contains {
		require.Equal(t, val, response.Teams, "teams comparison failed")
	} else {
		require.Equal(t, []string{}, response.Teams, "teams comparison failed")
	}

	if val, contains := overwrite["isSeen"]; contains {
		require.Equal(t, val.(bool), response.IsSeen, "isseen comparison failed")
	} else {
		require.Equal(t, false, response.IsSeen, "isseen comparison failed")
	}

	if val, contains := overwrite["acknowledged"]; contains {
		require.Equal(t, val.(bool), response.Acknowledged, "acknowledged comparison failed")
	} else {
		require.Equal(t, false, response.Acknowledged, "acknowledged comparison failed")
	}

	require.True(t, response.CreatedAt > 0, "CreatedAt should be greater than 0")
	require.True(t, response.UpdatedAt > 0, "UpdatedAt should be greater than 0")

	if val, contains := overwrite["count"]; contains {
		require.Equal(t, val.(int), response.Count, "count comparison failed")
	} else {
		require.Equal(t, 1, response.Count, "count comparison failed")
	}
	require.Equal(t, map[string]string{"prop1": "val1", "prop2": "val2"}, response.Details, "Details comparison failed")
}

func createAlert(t *testing.T, suffix string) string {
	req := alerts.CreateAlertRequest{
		Message:     "message" + suffix,
		Description: "description" + suffix,
		Source:      "source" + suffix,
		Entity:      "entity" + suffix,
		Alias:       "alert" + suffix,
		Note:        "Created for testing purposes" + suffix,
		User:        "user1",
		Actions:     testCfg.Alert.Actions,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Recipients:  []string{"user1"},
		Details:     map[string]string{"prop1": "val1", "prop2": "val2"},
	}
	response, alertErr := cli.Create(req)

	require.Nil(t, alertErr)
	require.NotNil(t, response)
	id := response.AlertId
	require.NotNil(t, id)
	return id
}

func deleteAllAlerts(t *testing.T) {
	for {
		listReq := alerts.ListAlertsRequest{}
		listResp, listErr := cli.List(listReq)

		require.Nil(t, listErr)
		require.NotNil(t, listResp)
		alertList := listResp.Alerts
		if len(alertList) == 0 {
			break
		} else {
			for i := 0; i < len(alertList); i++ {
				alert := alertList[i]
				delreq := alerts.DeleteAlertRequest{Id: alert.Id}
				delResp, alertErr := cli.Delete(delreq)

				require.Nil(t, alertErr)
				require.NotNil(t, delResp)
				require.Equal(t, 200, delResp.Code, "Response Code should be 200")
			}
		}
	}
}
