// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package api

import (
	"fmt"
	"net/http"
)

type TagsService struct {
	client *Client
}

type TagListObject []TagObject

type TagObject struct {
	// Digest        string    `json:"digest"`
	Name          string    `json:"name,omitempty"`
	// Size          int       `json:"size"`
	// Architecture  string    `json:"architecture"`
	// Os            string    `json:"os"`
	// OsVersion     string    `json:"os.version"`
	// DockerVersion string    `json:"docker_version"`
	// Author        string    `json:"author,omitempty"`
	// Created       time.Time `json:"created,omitempty"`
	//Config        struct {
	//	Labels struct {
	//		Maintainer string `json:"maintainer"`
	//	} `json:"labels"`
	//} `json:"config"`
	// Immutable bool          `json:"immutable"`
	// Signature interface{}   `json:"signature"`
	// Labels    []interface{} `json:"labels"`
	// PushTime  time.Time     `json:"push_time,omitempty"`
	// PullTime  time.Time     `json:"pull_time,omitempty"`
}

func (p *TagsService) List(name string) (v *TagListObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/repositories/%v/tags", p.client.endpoint, name)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	p.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(TagListObject)
	resp, err = p.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

