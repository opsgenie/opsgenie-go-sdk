package test

import (
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strings"
	"time"
)

type HeartbeatTestSuite struct {
	suite.Suite
}

func (suite *HeartbeatTestSuite) SetupTest() {
	deleteAllHeartbeats(suite.T())
}

// Heartbeat tests
func  (suite *HeartbeatTestSuite) TestAddHeartbeat() {
	t := suite.T()
	suffix := time.Now().String()
	id := createHeartbeat(t, "hb" + suffix)

	//get heartbeat with id
	getReq := hb.GetHeartbeatRequest{Id: id}
	getResp, hbErr := hbCli.Get(getReq)

	require.Nil(t, hbErr)
	require.NotNil(t, getResp)
	assertHeartbeat(t, id, getResp, suffix)
	t.Log("[OK] get heartbeat with id successful")

	//get heartbeat with name
	getReq = hb.GetHeartbeatRequest{Name: "hb" + suffix}
	getResp, hbErr = hbCli.Get(getReq)

	require.Nil(t, hbErr)
	require.NotNil(t, getResp)
	assertHeartbeat(t, id, getResp, suffix)
	t.Log("[OK] get heartbeat with name successful")
}

func  (suite *HeartbeatTestSuite) TestUpdateHeartbeat() {
	t := suite.T()
	suffix := time.Now().String()
	id := createHeartbeat(t, "hb" + suffix)

	enabled := false
	updateReq := hb.UpdateHeartbeatRequest{Id: id, Name: "updated heartbeat" + suffix, Description: "Some description", Interval: 20, IntervalUnit: "minutes", Enabled: &enabled}
	updateResp, updateErr := hbCli.Update(updateReq)

	require.Nil(t, updateErr)
	require.NotNil(t, updateResp)
	require.Equal(t, 200, updateResp.Code, "Response Code should be 200")

	//get heartbeat with id
	getReq := hb.GetHeartbeatRequest{Id: id}
	getResp, hbErr := hbCli.Get(getReq)

	require.Nil(t, hbErr)
	require.NotNil(t, getResp)

	require.Equal(t, id, getResp.Id, "id comparison failed")
	require.Equal(t, "updated heartbeat" + suffix, getResp.Name, "name comparison failed")
	require.Equal(t, uint64(0), getResp.LastHeartbeat, "last heartbeat time comparison failed")
	require.Equal(t, "Expired", getResp.Status, "status comparison failed")
	require.Equal(t, 20, getResp.Interval, "interval comparison failed")
	require.Equal(t, "minutes", getResp.IntervalUnit, "interval unit comparison failed")
	require.Equal(t, "Some description", getResp.Description, "Description comparison failed")
	require.False(t, getResp.Enabled, "Heartbeat enabled should be false")

	t.Log("[OK] heartbeat updated")
}

func  (suite *HeartbeatTestSuite) TestEnableHeartbeat() {
	t := suite.T()
	suffix := time.Now().String()
	enabled := false
	req := hb.AddHeartbeatRequest{Name: "hb"+suffix, Enabled: &enabled}
	response, hbErr := hbCli.Add(req)

	require.Nil(t, hbErr)
	require.NotNil(t, response)
	id := response.Id
	require.NotNil(t, id)

	enableReq := hb.EnableHeartbeatRequest{Id: id}
	enableResp, enableErr := hbCli.Enable( enableReq )

	require.Nil(t, enableErr)
	require.NotNil(t, enableResp)
	require.Equal(t, 200, enableResp.Code, "Response Code should be 200")

	getReq := hb.GetHeartbeatRequest{Id: id}
	getResp, hbErr := hbCli.Get(getReq)

	require.Nil(t, hbErr)
	require.NotNil(t, getResp)
	require.True(t, getResp.Enabled, "Heartbeat enabled should be true")
	t.Log("[OK] heartbeat enabled")
}

func  (suite *HeartbeatTestSuite) TestDisableHeartbeat() {
	t := suite.T()
	suffix := time.Now().String()
	enabled := true
	req := hb.AddHeartbeatRequest{Name: "hb"+suffix, Enabled: &enabled}
	response, hbErr := hbCli.Add(req)

	require.Nil(t, hbErr)
	require.NotNil(t, response)
	id := response.Id
	require.NotNil(t, id)

	disableReq := hb.DisableHeartbeatRequest{Id: id}
	disableResp, disableErr := hbCli.Disable( disableReq )

	require.Nil(t, disableErr)
	require.NotNil(t, disableResp)
	require.Equal(t, 200, disableResp.Code, "Response Code should be 200")

	getReq := hb.GetHeartbeatRequest{Id: id}
	getResp, hbErr := hbCli.Get(getReq)

	require.Nil(t, hbErr)
	require.NotNil(t, getResp)
	require.False(t, getResp.Enabled, "Heartbeat enabled should be false")
	t.Log("[OK] heartbeat disabled")
}

func  (suite *HeartbeatTestSuite) TestSendHeartbeat() {
	t := suite.T()
	suffix := time.Now().String()
	id := createHeartbeat(t, "hb" + suffix)

	sendReq := hb.SendHeartbeatRequest{Name: "hb"+suffix}
	sendResp, sendErr := hbCli.Send(sendReq)

	require.Nil(t, sendErr)
	require.NotNil(t, sendResp)
	require.True(t, sendResp.Heartbeat > uint64(0), "Heartbeat should greater than 0")

	getReq := hb.GetHeartbeatRequest{Id: id}
	getResp, hbErr := hbCli.Get(getReq)

	require.Nil(t, hbErr)
	require.NotNil(t, getResp)
	require.True(t, getResp.LastHeartbeat > uint64(0), "last heartbeat time comparison failed")
	require.Equal(t, "Active" , getResp.Status, "heartbeat status comparison failed")
	t.Log("[OK] heartbeat sent successfully")
}

func  (suite *HeartbeatTestSuite) TestDeleteHeartbeat() {
	t := suite.T()
	suffix := time.Now().String()
	id := createHeartbeat(t, "hb" + suffix)

	deleteReq := hb.DeleteHeartbeatRequest{Id: id}
	deleteResp, deleteErr := hbCli.Delete( deleteReq )

	require.Nil(t, deleteErr)
	require.NotNil(t, deleteResp)
	require.Equal(t, 200, deleteResp.Code, "Response Code should be 200")

	getReq := hb.GetHeartbeatRequest{Id: id }
	getResp, getErr := hbCli.Get(getReq)

	require.NotNil(t, getErr)
	require.Nil(t, getResp)
	require.True(t, strings.Contains(getErr.Error(), "Heartbeat with id ["+ id +"] does not exist"), "Error should contains ")
	t.Log("[OK] heartbeat deleted successfully")
}

func (suite *HeartbeatTestSuite) TestListHeartbeats() {
	t := suite.T()
	listReq := hb.ListHeartbeatsRequest{}
	listResp, listErr := hbCli.List(listReq)

	require.Nil(t, listErr)
	require.NotNil(t, listResp)

	require.Equal(t, 0, len(listResp.Heartbeats))

	id := createHeartbeat(t, "hb1")

	listReq = hb.ListHeartbeatsRequest{}
	listResp, listErr = hbCli.List(listReq)

	require.Nil(t, listErr)
	require.NotNil(t, listResp)
	beats := listResp.Heartbeats
	require.Equal(t, 1, len(beats))

	found := false

	for _, beat := range(beats) {
		if beat.Id == id {
			found = true
			require.Equal(t, "hb1", beat.Name, "name comparison failed")
			require.Equal(t, uint64(0), beat.LastHeartbeat, "last heartbeat time comparison failed")
			require.Equal(t, "Expired", beat.Status, "status comparison failed")
		}
	}
	require.True(t, found, "newly created id could not be found in heartbeat list")
	t.Log("[OK] heartbeats list succesfully")
}

func deleteAllHeartbeats(t *testing.T){
	for{
		listReq := hb.ListHeartbeatsRequest{}
		listResp, listErr := hbCli.List(listReq)

		require.Nil(t, listErr)
		require.NotNil(t, listResp)
		hbList := listResp.Heartbeats
		if len(hbList) == 0 {
			t.Log("All heartbeats deleted")
			break
		}else{
			for i:=0;i<len(hbList); i++{
				heartbeat := hbList[i]
				delreq := hb.DeleteHeartbeatRequest{Id: heartbeat.Id}
				delResp, delErr := hbCli.Delete(delreq)

				require.Nil(t, delErr)
				require.NotNil(t, delResp)
				require.Equal(t, 200, delResp.Code, "Response Code should be 200")
			}
		}
	}
}

func assertHeartbeat(t *testing.T, id string, getResp *hb.GetHeartbeatResponse, suffix string){
	require.Equal(t, id, getResp.Id, "id comparison failed")
	require.Equal(t, "hb" + suffix, getResp.Name, "name comparison failed")
	require.Equal(t, uint64(0), getResp.LastHeartbeat, "last heartbeat time comparison failed")
	require.Equal(t, "Expired", getResp.Status, "status comparison failed")
	require.Equal(t, 10, getResp.Interval, "interval comparison failed")
	require.Equal(t, "hours", getResp.IntervalUnit, "interval unit comparison failed")
	require.Equal(t, "descr", getResp.Description, "Description comparison failed")
	require.True(t, getResp.Enabled, "Heartbeat enabled should be true")
}

func createHeartbeat(t *testing.T, name string) string{
	req := hb.AddHeartbeatRequest{Name: name, Description: "descr", Interval: 10, IntervalUnit: "hours"}
	response, hbErr := hbCli.Add(req)
	require.Nil(t, hbErr)
	require.NotNil(t, response)
	id := response.Id
	require.NotNil(t, id)
	return id
}