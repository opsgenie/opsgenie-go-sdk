package userv2

import (
	"errors"
	"net/url"
	"strconv"
)

const (
	// The list of user fields.
	UsernameQueryField    = "username"
	FullNameQueryField    = "fullName"
	BlockedQueryField     = "blocked"
	VerifiedQueryField    = "verified"
	RoleQueryField        = "role"
	LocaleQueryField      = "locale"
	TimeZoneQueryField    = "timeZone"
	UserAddressQueryField = "userAddress"
	CreatedAtQueryField   = "createdAt"
)

// ListUsersRequest is a request for getting user list.
type ListUsersRequest struct {
	Limit  int
	Offset int
	Sort   Sort
	Order  Order
	Query  map[string]string
	ApiKey string
}

// GenerateUrl generates API url for getting user list.
func (request *ListUsersRequest) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/users/"

	params := url.Values{}
	params.Set("offset", strconv.Itoa(request.Offset))

	if request.Limit > 0 {
		params.Set("limit", strconv.Itoa(request.Limit))
	}

	if len(request.Sort) > 0 {
		if request.Sort == UsernameSortField || request.Sort == FullNameSortField || request.Sort == InsertedAtSortField {
			params.Set("sort", string(request.Sort))
		} else {
			return "", nil, errors.New("unavailable field to use in sorting, id should be one of 'username', 'fullName' or 'insertedAt'")
		}
	}

	if len(request.Order) > 0 {
		if request.Order == AscSortType || request.Order == DescSortType {
			params.Set("order", string(request.Order))
		} else {
			return "", nil, errors.New("unavailable direction of sorting, it should be 'asc' or 'desc'")
		}
	}

	if len(request.Query) > 0 {
		for key, value := range request.Query {
			if key == UsernameQueryField ||
				key == FullNameQueryField ||
				key == BlockedQueryField ||
				key == VerifiedQueryField ||
				key == RoleQueryField ||
				key == LocaleQueryField ||
				key == TimeZoneQueryField ||
				key == UserAddressQueryField ||
				key == CreatedAtQueryField {
				params.Set(key, value)
			} else {
				return "", nil, errors.New("unavailable query field, it should be one of 'username', 'fullName', 'blocked', 'verified', 'role', 'locale', 'timeZone', 'userAddress' or 'createdAt'")
			}
		}
	}

	return baseUrl, params, nil
}

// GetApiKey returns api key.
func (r *ListUsersRequest) GetApiKey() string {
	return r.ApiKey
}
