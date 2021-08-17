// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package api

import (
	"fmt"
	"net/http"
)

type RepositoriesService struct {
	client *Client
}

type RepositoriesListObject []RepositoriesObject

type RepositoriesObject struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	ProjectID    int           `json:"project_id"`
	Description  string        `json:"description"`
	PullCount    int           `json:"pull_count"`
	StarCount    int           `json:"star_count"`
	TagsCount    int           `json:"tags_count"`
	// Labels       []interface{} `json:"labels"`
	// CreationTime time.Time     `json:"creation_time"`
	//UpdateTime   time.Time     `json:"update_time"`
}

func (p *RepositoriesService) List(id int) (v *RepositoriesListObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/repositories?project_id=%v&page=1&page_size=2000", p.client.endpoint, id)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	p.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RepositoriesListObject)
	resp, err = p.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

