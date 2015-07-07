package test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/franela/goreq"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	itg "github.com/opsgenie/opsgenie-go-sdk/integration"
	"github.com/opsgenie/opsgenie-go-sdk/policy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	yaml "gopkg.in/yaml.v2"
)

type ClientTestConfig struct {
	Alert struct {
		ApiKey  string   `yaml:"apiKey"`
		User    string   `yaml:"user"`
		Team    string   `yaml:"team"`
		Actions []string `yaml:"actions"`
	} `yaml:"alert"`

	OpsGenieApiUrl string `yaml:"opsgenie.api.url"`
}

type EntityNames struct {
	Integration string `json:"integration"`
	Team        string `json:"team"`
	Policy      string `json:"policy"`
}

// common globals
var cli *ogcli.OpsGenieAlertClient
var CONFIG_FILE_NAME string = "client_at_test_cfg.yaml"
var testCfg ClientTestConfig
var hbCli *ogcli.OpsGenieHeartbeatClient
var itgCli *ogcli.OpsGenieIntegrationClient
var policyCli *ogcli.OpsGeniePolicyClient
var entityNames EntityNames

func TestEnableDisableIntegration(t *testing.T) {
	disableReq := itg.DisableIntegrationRequest{Name: entityNames.Integration}
	disableResp, disableErr := itgCli.Disable(disableReq)

	require.Nil(t, disableErr)
	require.NotNil(t, disableResp)
	require.Equal(t, 200, disableResp.Code, "Response Code should be 200")
	require.Equal(t, "success", disableResp.Status, "Response Code should be 200")
	t.Log("[OK] integration disabled")

	enableReq := itg.EnableIntegrationRequest{Name: entityNames.Integration}
	enableResp, enableErr := itgCli.Enable(enableReq)

	require.Nil(t, enableErr)
	require.NotNil(t, enableResp)
	require.Equal(t, 200, enableResp.Code, "Response Code should be 200")
	require.Equal(t, "success", enableResp.Status, "Response Code should be 200")
	t.Log("[OK] integration enabled")
}

func TestEnableDisablePolicy(t *testing.T) {
	disableReq := policy.DisablePolicyRequest{Name: entityNames.Policy}
	disableResp, disableErr := policyCli.Disable(disableReq)
	require.Nil(t, disableErr)
	require.NotNil(t, disableResp)
	require.Equal(t, 200, disableResp.Code, "Response Code should be 200")
	t.Log("[OK] policy disabled")

	enableReq := policy.EnablePolicyRequest{Name: entityNames.Policy}
	enableResp, enableErr := policyCli.Enable(enableReq)

	require.Nil(t, enableErr)
	require.NotNil(t, enableResp)
	require.Equal(t, 200, enableResp.Code, "Response Code should be 200")
	t.Log("[OK] policy enabled")
}

func TestAlertClientSuite(t *testing.T) {
	suite.Run(t, new(AlertTestSuite))
}

func TestHeartbeatClientSuite(t *testing.T) {
	suite.Run(t, new(HeartbeatTestSuite))
}

// utility function
func readSettingsFromConfigFile() error {
	cfgData, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		return errors.New("Can not read from the configuration file: " + err.Error())
	}
	err = yaml.Unmarshal(cfgData, &testCfg)
	if err != nil {
		return errors.New("Can not parse the configuration file: " + err.Error())
	}
	return nil
}

// setup the test suite
func TestMain(m *testing.M) {
	// read the settings
	err := readSettingsFromConfigFile()
	if err != nil {
		panic(err)
	}

	// create an opsgenie client
	opsGenieClient := new(ogcli.OpsGenieClient)
	opsGenieClient.SetApiKey(testCfg.Alert.ApiKey)
	opsGenieClient.SetOpsGenieApiUrl(testCfg.OpsGenieApiUrl)

	req := goreq.Request{Method: "POST", Uri: opsGenieClient.GetOpsGenieApiUrl() + "/v1/json/sdkSetup", Body: map[string]string{"apiKey": opsGenieClient.GetApiKey()}}
	resp, err := req.Do()
	if err != nil {
		fmt.Printf("Could not send request to create test team, integration, policy; %s\n", err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
		if err = resp.Body.FromJsonTo(&entityNames); err != nil {
			fmt.Printf("Server response for sdkSetup can not be parsed,  %s\n", err.Error())
		}
	}

	var cliErr error
	cli, cliErr = opsGenieClient.Alert()

	if cliErr != nil {
		panic(cliErr.Error())
	}

	hbCli, cliErr = opsGenieClient.Heartbeat()

	if cliErr != nil {
		panic(cliErr)
	}

	itgCli, cliErr = opsGenieClient.Integration()

	if cliErr != nil {
		panic(cliErr)
	}

	policyCli, cliErr = opsGenieClient.Policy()

	if cliErr != nil {
		panic(cliErr)
	}

	os.Exit(m.Run())
}
