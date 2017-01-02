/*
Copyright 2016. All rights reserved.
Use of this source code is governed by a Apache Software
license that can be found in the LICENSE file.
*/

//Package group provides requests and response structures to achieve Group API actions.
package group

// CreateGroupRequest provides necessary parameter structure for creating a group
type CreateGroupRequest struct {
	APIKey string `json:"apiKey,omitempty"`
	Name   string `json:"name,omitempty"`
	Users []string `json:"users,omitempty"`
}

// UpdateGroupRequest provides necessary parameter structure for updating a group
type UpdateGroupRequest struct {
	APIKey string `json:"apiKey,omitempty"`
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Users []string `json:"users,omitempty"`
}

// DeleteGroupRequest provides necessary parameter structure for deleting a group
type DeleteGroupRequest struct {
	APIKey string `url:"apiKey,omitempty"`
	Id     string `url:"id,omitempty"`
	Name   string `url:"name,omitempty"`
}

// GetGroupRequest provides necessary parameter structure for requesting information about a group
type GetGroupRequest struct {
	APIKey string `url:"apiKey,omitempty"`
	Id     string `url:"id,omitempty"`
	Name   string `url:"name,omitempty"`
}

// ListGroupsRequest provides necessary parameter structure for listing groups
type ListGroupsRequest struct {
	APIKey string `url:"apiKey,omitempty"`
}
