// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package api

import (
	"fmt"
	"net/http"
	"time"
)

type ProjectService struct {
	client *Client
}

type ProjectListObject []ProjectObject

type ProjectMetadata struct {
	Public string `json:"public"`
	RetentionID string `json:"retention_id,omitempty"`
	AutoScan             string `json:"auto_scan,omitempty"`
	EnableContentTrust   string `json:"enable_content_trust,omitempty"`
	PreventVul           string `json:"prevent_vul,omitempty"`
	ReuseSysCveWhitelist string `json:"reuse_sys_cve_whitelist,omitempty"`
	Severity             string `json:"severity,omitempty"`
}

type ProjectCveWhitelist struct {
	ID           int         `json:"id,omitempty"`
	ProjectID    int         `json:"project_id,omitempty"`
	Items        interface{} `json:"items,omitempty"`
	CreationTime time.Time   `json:"creation_time,omitempty"`
	UpdateTime   time.Time   `json:"update_time,omitempty"`
}

type ProjectObject struct {
	ProjectID          int         `json:"project_id"`
	OwnerID            int         `json:"owner_id"`
	Name               string      `json:"name"`
	// CreationTime       time.Time   `json:"creation_time"`
	// UpdateTime         time.Time   `json:"update_time"`
	Deleted            bool        `json:"deleted"`
	// OwnerName          string      `json:"owner_name"`
	// CurrentUserRoleID  int         `json:"current_user_role_id"`
	//CurrentUserRoleIds interface{} `json:"current_user_role_ids"`
	RepoCount          int         `json:"repo_count"`
	ChartCount         int         `json:"chart_count"`
	Metadata           ProjectMetadata `json:"metadata,omitempty"`
	//CveWhitelist ProjectCveWhitelist `json:"cve_whitelist,omitempty"`
}

func (p *ProjectService) List() (v *ProjectListObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/projects?page=1&page_size=1000", p.client.endpoint)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	p.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectListObject)
	resp, err = p.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}