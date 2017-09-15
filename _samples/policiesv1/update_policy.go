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
		Type: policyv1.MatchAnyConditionFilterType,
		Conditions: []policyv1.Condition{
			{
				Field:         policyv1.ExtraPropertiesField,
				Key:           "first_property",
				Not:           false,
				Operation:     policyv1.EqualsOperationType,
				ExpectedValue: "first_value",
			},
			{
				Field:         policyv1.ExtraPropertiesField,
				Key:           "second_property",
				Not:           true,
				Operation:     policyv1.StartsWithOperationType,
				ExpectedValue: "second_value",
			},
		},
	}

	startHour := 5
	endHour := 9
	startMin := 10
	endMin := 45

	timeRestrictions := policyv1.TimeRestrictions{
		Type: policyv1.TimeOfDayTimeRestrictionType,
		Restriction: policyv1.Restriction{
			StartHour: &startHour,
			EndHour:   &endHour,
			StartMin:  &startMin,
			EndMin:    &endMin,
		},
	}

	timeAmount := 4
	duration := policyv1.Duration{
		TimeAmount: &timeAmount,
		TimeUnit:   policyv1.HoursTimeUnit,
	}

	updatePolicyRequest := policyv1.UpdatePolicyRequest{
		Identifier: &policyv1.Identifier{ID: "1"},
		Name:              "updated_policy_name",
		PolicyDescription: "it is updated policy",
		Enabled:           false,
		Filter:            filter,
		TimeRestrictions:  &timeRestrictions,
		Type:              policyv1.NotificationSuppressPolicyType,
		Duration:          duration,
	}

	response, err := policyCli.Update(updatePolicyRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", response.Status)
}
