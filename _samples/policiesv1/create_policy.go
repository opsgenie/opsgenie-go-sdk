package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/policyv1"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	policyCli, err := cli.PolicyV1()
	if err != nil {
		panic(err)
	}

	filter := policyv1.Filter{
		Type: policyv1.MatchAllConditionsFilterType,
		Conditions: []policyv1.Condition{
			{
				Field:         policyv1.ExtraPropertiesField,
				Key:           "test_key",
				Not:           false,
				Operation:     policyv1.ContainsOpertionType,
				ExpectedValue: "test_value",
			},
		},
	}

	startDay := policyv1.Monday
	endDay := policyv1.Wednesday
	startHour := 3
	endHour := 15
	startMin := 5
	endMin := 30

	timeRestrictions := policyv1.TimeRestrictions{
		Type: policyv1.WeekdayAndTimeOfDayTimeRestrictionType,
		Restrictions: []policyv1.Restriction{
			{
				StartDay:  startDay,
				EndDay:    endDay,
				StartHour: &startHour,
				EndHour:   &endHour,
				StartMin:  &startMin,
				EndMin:    &endMin,
			},
		},
	}

	timeAmount := 1
	duration := policyv1.Duration{
		TimeAmount: &timeAmount,
		TimeUnit:   policyv1.MinutesTimeUnit,
	}

	createPolicyRequest := policyv1.CreateNotificationSuppressPolicyRequest{
		CreatePolicyRequest: policyv1.CreatePolicyRequest{
			Name:              "test_policy_name",
			PolicyDescription: "it is test policy",
			Enabled:           true,
			Filter:            filter,
			TimeRestrictions:  &timeRestrictions,
			Type:              policyv1.NotificationSuppressPolicyType,
		},
	}

	response, err := policyCli.CreateNotificationSuppressPolicy(createPolicyRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"ID: %s\nName: %s\nType: %s\nEnabled: %t\n",
		response.Policy.ID,
		response.Policy.Name,
		response.Policy.Type,
		response.Policy.Enabled,
	)
}
