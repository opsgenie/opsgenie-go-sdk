package test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	"github.com/stretchr/testify/require"
)

type renotifyWaitAction struct {
	ID string
}

func (act renotifyWaitAction) check(t *testing.T) bool {
	listLogsReq := alerts.ListAlertLogsRequest{ID: act.ID}
	listLogsResponse, alertErr := cli.ListLogs(listLogsReq)

	require.Nil(t, alertErr)
	require.NotNil(t, listLogsResponse)
	logs := listLogsResponse.Logs
	for i := 0; i < len(logs); i++ {
		var alertLog = logs[i]
		if strings.Contains(alertLog.Log, "Renotifying") {
			return true
		}
	}
	return false
}

func (suite *alertTestSuite) TestRenotifyAlert() {
	t := suite.T()
	suffix := time.Now().String()
	user := "owneruser"
	id := createAlert(t, suffix)

	renotifyReq := alerts.RenotifyAlertRequest{ID: id, User: user}
	renotifyResp, alertErr := cli.Renotify(renotifyReq)

	require.Nil(t, alertErr)
	require.NotNil(t, renotifyResp)
	require.Equal(t, 200, renotifyResp.Code, "Response Code should be 200")

	require.True(t, waitFor(t, renotifyWaitAction{ID: id}), "Alert logs should contain renotifying log.")

	//renotify with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)

	renotifyReq = alerts.RenotifyAlertRequest{Alias: "alert" + suffix, User: user}
	renotifyResp, alertErr = cli.Renotify(renotifyReq)

	require.Nil(t, alertErr)
	require.NotNil(t, renotifyResp)
	require.Equal(t, 200, renotifyResp.Code, "Response Code should be 200")

	require.True(t, waitFor(t, renotifyWaitAction{ID: id}), "Alert logs should contain renotifying log.")
	t.Log("[OK] renotified successfully")
}

type listAlertLogsWaitAction struct {
	ID string
}

func (act listAlertLogsWaitAction) check(t *testing.T) bool {
	listLogsReq := alerts.ListAlertLogsRequest{ID: act.ID}
	listLogsResponse, alertErr := cli.ListLogs(listLogsReq)

	require.Nil(t, alertErr)
	require.NotNil(t, listLogsResponse)
	logs := listLogsResponse.Logs

	for i := 0; i < len(logs); i++ {
		var alertLog = logs[i]
		if strings.Contains(alertLog.Log, "Alert created") {
			require.True(t, alertLog.CreatedAt > 0)
			require.Equal(t, "System", alertLog.Owner)
			require.Equal(t, "system", alertLog.LogType)
			return true
		}
	}
	return false
}

func (suite *alertTestSuite) TestListAlertLogs() {
	t := suite.T()
	suffix := time.Now().String()
	req := alerts.CreateAlertRequest{Message: "message" + suffix, Recipients: []string{"user1"}}
	response, alertErr := cli.Create(req)

	require.Nil(t, alertErr)
	require.NotNil(t, response)
	id := response.AlertID
	require.NotNil(t, id)

	require.True(t, waitFor(t, listAlertLogsWaitAction{ID: id}), "Alert logs should contain alert created log.")
	t.Log("[OK] listed alert logs successfully")
}

type listAlertNotesWaitAction struct {
	ID string
}

func (act listAlertNotesWaitAction) check(t *testing.T) bool {
	listNotesReq := alerts.ListAlertNotesRequest{ID: act.ID}
	listNotesResponse, alertErr := cli.ListNotes(listNotesReq)

	require.Nil(t, alertErr)
	require.NotNil(t, listNotesResponse)

	notes := listNotesResponse.Notes
	for i := 0; i < len(notes); i++ {
		var alertNote = notes[i]
		if strings.Contains(alertNote.Note, "Alert note") {
			require.True(t, alertNote.CreatedAt > 0)
			require.Equal(t, "actionuser", alertNote.Owner)
			return true
		}
	}
	return false
}

func (suite *alertTestSuite) TestAddNoteAndListNotes() {
	t := suite.T()
	suffix := time.Now().String()
	id := createAlert(t, suffix)
	addNoteReq := alerts.AddNoteAlertRequest{ID: id, Note: "Alert note", User: "actionuser"}
	addNoteResp, alertErr := cli.AddNote(addNoteReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addNoteResp)
	require.Equal(t, 200, addNoteResp.Code, "Response code should be 200")

	t.Log("[OK] note added with to alert")

	require.True(t, waitFor(t, listAlertNotesWaitAction{ID: id}), "Alert notes should contain the recently added note.")
	t.Log("[OK] listed notes successfully")
}

func (suite *alertTestSuite) TestAddTeam() {
	t := suite.T()

	suffix := time.Now().String()
	id := createAlert(t, suffix)
	addTeamReq := alerts.AddTeamAlertRequest{ID: id, Team: entityNames.Team, User: "somebody", Note: "note1"}
	addTeamResponse, alertErr := cli.AddTeam(addTeamReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addTeamResponse)
	require.Equal(t, 200, addTeamResponse.Code, "Response code should be 200")

	getAlertReq := alerts.GetAlertRequest{ID: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"teams": []string{entityNames.Team}})
	t.Log("[OK] team added with id successfully")

	//Add team with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	addTeamReq = alerts.AddTeamAlertRequest{Alias: "alert" + suffix, Team: entityNames.Team}
	addTeamResponse, alertErr = cli.AddTeam(addTeamReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addTeamResponse)
	require.Equal(t, 200, addTeamResponse.Code, "Response code should be 200")

	getAlertReq = alerts.GetAlertRequest{ID: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"teams": []string{entityNames.Team}})
	t.Log("[OK] team added with alias successfully")
}

func (suite *alertTestSuite) TestAddTags() {
	t := suite.T()

	suffix := time.Now().String()
	id := createAlert(t, suffix)
	addTagsReq := alerts.AddTagsAlertRequest{ID: id, Tags: []string{"newlyAddedTag1", "newlyAddedTag2"}, User: "somebody", Note: "note1"}
	addTagsResp, alertErr := cli.AddTags(addTagsReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addTagsResp)
	require.Equal(t, 200, addTagsResp.Code, "Response code should be 200")

	getAlertReq := alerts.GetAlertRequest{ID: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"tags": []string{"newlyAddedTag1", "newlyAddedTag2", "tag1", "tag2"}})
	t.Log("[OK] tags added with id successfully")

	//Add tags with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	addTagsReq = alerts.AddTagsAlertRequest{Alias: "alert" + suffix, Tags: []string{"newlyAddedTag1"}}
	addTagsResp, alertErr = cli.AddTags(addTagsReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addTagsResp)
	require.Equal(t, 200, addTagsResp.Code, "Response code should be 200")

	getAlertReq = alerts.GetAlertRequest{ID: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"tags": []string{"newlyAddedTag1", "tag1", "tag2"}})

	t.Log("[OK] tags added with alias successfully")
}

func (suite *alertTestSuite) TestAssignOwner() {
	t := suite.T()

	suffix := time.Now().String()
	id := createAlert(t, suffix)
	assignOwnerReq := alerts.AssignOwnerAlertRequest{ID: id, Owner: "owneruser", User: "actionuser"}
	assignOwnerResponse, alertErr := cli.AssignOwner(assignOwnerReq)

	require.Nil(t, alertErr)
	require.NotNil(t, assignOwnerResponse)
	require.Equal(t, 200, assignOwnerResponse.Code, "Response code should be 200")

	getAlertReq := alerts.GetAlertRequest{ID: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"owner": "owneruser"})
	t.Log("[OK] assigned with id successfully")

	//assign with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	assignOwnerReq = alerts.AssignOwnerAlertRequest{Alias: "alert" + suffix, Owner: "owneruser", User: "actionuser"}
	assignOwnerResponse, alertErr = cli.AssignOwner(assignOwnerReq)

	require.Nil(t, alertErr)
	require.NotNil(t, assignOwnerResponse)
	require.Equal(t, 200, assignOwnerResponse.Code, "Response code should be 200")

	getAlertReq = alerts.GetAlertRequest{ID: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"owner": "owneruser"})
	t.Log("[OK] assigned with alias successfully")
}

type executeActionWaitAction struct {
	ID string
}

func (act executeActionWaitAction) check(t *testing.T) bool {
	listLogsReq := alerts.ListAlertLogsRequest{ID: act.ID}
	listLogsResponse, alertErr := cli.ListLogs(listLogsReq)

	require.Nil(t, alertErr)
	require.NotNil(t, listLogsResponse)
	logs := listLogsResponse.Logs

	for i := 0; i < len(logs); i++ {
		var alertLog = logs[i]
		if strings.Contains(alertLog.Log, testCfg.Alert.Actions[0]) {
			require.True(t, alertLog.CreatedAt > 0)
			require.Equal(t, "System", alertLog.Owner)
			require.Equal(t, "alertAction", alertLog.LogType)
			return true
		}
	}

	return false
}

func (suite *alertTestSuite) TestExecuteAction() {
	t := suite.T()

	suffix := time.Now().String()
	id := createAlert(t, suffix)
	execActionReq := alerts.ExecuteActionAlertRequest{ID: id,
		Action: testCfg.Alert.Actions[0],
		Note:   fmt.Sprintf("Action <b>%s</b> executed by the Go API", testCfg.Alert.Actions[0]),
	}
	execActionResponse, alertErr := cli.ExecuteAction(execActionReq)

	require.Nil(t, alertErr)
	require.NotNil(t, execActionResponse)
	require.True(t, strings.Contains(execActionResponse.Result, "Initiated ["+testCfg.Alert.Actions[0]+"] action"), "result should contain")
	require.Equal(t, 200, execActionResponse.Code, "Response code should be 200")

	require.True(t, waitFor(t, executeActionWaitAction{ID: id}), "Alert logs should contain action execution log.")
	t.Log("[OK] action " + testCfg.Alert.Actions[0] + " executed on alert")

	//execute action with alias
	execActionReq = alerts.ExecuteActionAlertRequest{Alias: "alert" + suffix,
		Action: testCfg.Alert.Actions[1],
		Note:   fmt.Sprintf("Action <b>%s</b> executed by the Go API", testCfg.Alert.Actions[1]),
	}
	execActionResponse, alertErr = cli.ExecuteAction(execActionReq)

	require.Nil(t, alertErr)
	require.NotNil(t, execActionResponse)
	require.True(t, strings.Contains(execActionResponse.Result, "Initiated ["+testCfg.Alert.Actions[1]+"] action"), "result should contain")
	require.Equal(t, 200, execActionResponse.Code, "Response code should be 200")

	require.True(t, waitFor(t, executeActionWaitAction{ID: id}), "Alert logs should contain action execution log.")
	t.Log("[OK] action " + testCfg.Alert.Actions[1] + " executed on alert")
}

func (suite *alertTestSuite) TestTakeOwnership() {
	t := suite.T()

	suffix := time.Now().String()
	id := createAlert(t, suffix)
	takeOwnershipReq := alerts.TakeOwnershipAlertRequest{ID: id, User: "user1", Note: "ownership note", Source: "go api"}

	takeOwnershipResp, alertErr := cli.TakeOwnership(takeOwnershipReq)

	require.Nil(t, alertErr)
	require.NotNil(t, takeOwnershipResp)
	require.Equal(t, 200, takeOwnershipResp.Code, "Response code should be 200")

	getAlertReq := alerts.GetAlertRequest{ID: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"owner": "user1", "isSeen": true})
	t.Log("[OK] take ownership successful")

	//take ownership with alias
	suffix = time.Now().String()
	id = createAlert(t, suffix)
	takeOwnershipReq = alerts.TakeOwnershipAlertRequest{ID: id, User: "user1", Note: "ownership note", Source: "go api"}

	takeOwnershipResp, alertErr = cli.TakeOwnership(takeOwnershipReq)

	require.Nil(t, alertErr)
	require.NotNil(t, takeOwnershipResp)
	require.Equal(t, 200, takeOwnershipResp.Code, "Response code should be 200")

	getAlertReq = alerts.GetAlertRequest{ID: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"owner": "user1", "isSeen": true})
	t.Log("[OK] take ownership successful")
}

func (suite *alertTestSuite) TestAddRecipient() {
	t := suite.T()

	suffix := time.Now().String()
	id := createAlert(t, suffix)
	addRecipientReq := alerts.AddRecipientAlertRequest{ID: id, User: "user1", Note: "add recipient note", Recipient: "newrecipient"}

	addrecipientResp, alertErr := cli.AddRecipient(addRecipientReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addrecipientResp)
	require.Equal(t, 200, addrecipientResp.Code, "Response code should be 200")

	getAlertReq := alerts.GetAlertRequest{ID: id}
	getResp, alertErr := cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"recipients": []string{"newrecipient", "user1"}})
	t.Log("[OK] add recipient successful")

	suffix = time.Now().String()
	id = createAlert(t, suffix)
	addRecipientReq = alerts.AddRecipientAlertRequest{Alias: "alert" + suffix, User: "user1", Note: "add recipient note", Recipient: "newrecipient"}

	addrecipientResp, alertErr = cli.AddRecipient(addRecipientReq)

	require.Nil(t, alertErr)
	require.NotNil(t, addrecipientResp)
	require.Equal(t, 200, addrecipientResp.Code, "Response code should be 200")

	getAlertReq = alerts.GetAlertRequest{ID: id}
	getResp, alertErr = cli.Get(getAlertReq)

	require.Nil(t, alertErr)
	require.NotNil(t, getResp)
	assertAlert(t, id, getResp, suffix, map[string]interface{}{"recipients": []string{"newrecipient", "user1"}})
	t.Log("[OK] add recipient with alias successful")
}

type listRecipientsWaitAction struct {
	ID string
}

func (act listRecipientsWaitAction) check(t *testing.T) bool {
	listRecipientsReq := alerts.ListAlertRecipientsRequest{ID: act.ID}
	listRecipientsResp, alertErr := cli.ListRecipients(listRecipientsReq)
	require.Nil(t, alertErr)
	require.NotNil(t, listRecipientsResp)

	if len(listRecipientsResp.Groups) != 0 {
		return false
	}

	if len(listRecipientsResp.Users) == 1 {
		recp := listRecipientsResp.Users[0]
		require.Equal(t, testCfg.Alert.User, recp.Username)
		require.True(t, recp.State != "", "recipient state should not be empty string")
		require.True(t, recp.StateChangedAt > 0, "recipient state changed at should be greater than 0")
		require.NotNil(t, recp.Method)
		return true
	}
	return false
}

func (suite *alertTestSuite) TestListAlertRecipients() {
	t := suite.T()

	suffix := time.Now().String()
	req := alerts.CreateAlertRequest{Message: "message" + suffix, Recipients: []string{testCfg.Alert.User}}
	response, alertErr := cli.Create(req)

	require.Nil(t, alertErr)
	require.NotNil(t, response)
	id := response.AlertID
	require.NotNil(t, id)

	require.True(t, waitFor(t, listRecipientsWaitAction{ID: id}), "List recipients failed.")
	t.Log("[OK] listed alert recipients successfully")
}

func (suite *alertTestSuite) TestAttachFile() {
	t := suite.T()

	suffix := time.Now().String()
	req := alerts.CreateAlertRequest{Message: "message" + suffix, Recipients: []string{testCfg.Alert.User}}
	response, alertErr := cli.Create(req)

	require.Nil(t, alertErr)
	require.NotNil(t, response)
	id := response.AlertID
	require.NotNil(t, id)

	fileName := "dummy.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	require.Nil(t, err)

	file.WriteString("dummytext")
	defer file.Close()

	attachFileReq := alerts.AttachFileAlertRequest{ID: id, Attachment: file, IndexFile: fileName, User: "someuser"}
	attachFileResp, attachFileErr := cli.AttachFile(attachFileReq)

	require.Nil(t, attachFileErr)
	require.NotNil(t, attachFileResp)
	require.Equal(t, 200, attachFileResp.Code, "Response Code should be 200")
	t.Log("[OK] attached file to alert successfully")

	defer os.Remove("dummy.txt")
}

func (suite *alertTestSuite) TestListAlerts() {
	t := suite.T()
	cnt := 10
	var ids [10]string
	for i := 0; i < cnt; i++ {
		suffix := strconv.Itoa(i)
		id := createAlert(t, suffix)
		ids[(cnt-1)-i] = id
		//		time.Sleep(500 * time.Millisecond)
	}
	t.Log(strconv.Itoa(cnt) + " alerts created.")

	listReq := alerts.ListAlertsRequest{}
	listResp, listErr := cli.List(listReq)

	require.Nil(t, listErr)
	require.NotNil(t, listResp)
	alertList := listResp.Alerts

	require.Equal(t, cnt, len(alertList), "alert count comparison failed")

	for i := 0; i < cnt; i++ {
		require.Equal(t, ids[i], alertList[i].ID)
	}

	listReq = alerts.ListAlertsRequest{CreatedAfter: alertList[0].CreatedAt}
	listResp, listErr = cli.List(listReq)

	require.Nil(t, listErr)
	require.NotNil(t, listResp)
	alertList = listResp.Alerts

	require.Equal(t, 0, len(alertList), "alert count comparison failed")

	listReq = alerts.ListAlertsRequest{Limit: 2, SortBy: "createdAt", Order: "asc"}
	listResp, listErr = cli.List(listReq)

	require.Nil(t, listErr)
	require.NotNil(t, listResp)
	alertList = listResp.Alerts

	require.Equal(t, 2, len(alertList), "alert count comparison failed")
	for i := 0; i < 2; i++ {
		require.Equal(t, ids[(cnt-1)-i], alertList[i].ID)
	}

	t.Log("Listed Alerts successfully")
}
