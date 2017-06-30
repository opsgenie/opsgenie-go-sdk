package userv2

import (
	"errors"
	"fmt"
	"net/url"
)

// Identifier defined the set of attributes for identification user.
type Identifier struct {
	ID       string `json:"-"`
	UserName string `json:"-"`
	Entity   Entity `json:"-"`
	Expand   Expand `json:"-"`
}

// GenerateUrl generates API url using specified attributes of identifier.
func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/users"

	if len(request.ID) > 0 {
		baseUrl += "/" + request.ID
	} else if len(request.UserName) > 0 {
		baseUrl += "/" + request.UserName
	} else {
		return "", nil, errors.New("username or id of the user should be provided")
	}

	if len(request.Entity) > 0 {
		if request.Entity == EscalationsEntity ||
			request.Entity == TeamsEntity ||
			request.Entity == ForwardingRulesEntity ||
			request.Entity == SchedulesEntity {

			baseUrl += "/" + string(request.Entity)
		} else {
			err := fmt.Errorf(
				"not available user entity, in should one of the %s, %s, %s, %s",
				EscalationsEntity,
				TeamsEntity,
				ForwardingRulesEntity,
				SchedulesEntity,
			)

			return "", nil, err
		}
	}

	params := url.Values{}
	if len(request.Expand) > 0 {
		if request.Expand == ContactExpandableField {
			params.Set("expand", string(request.Expand))
		} else {
			return "", nil, errors.New("the only expandable field for user api is 'contact'")
		}
	}

	return baseUrl, params, nil
}
