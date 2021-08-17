// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	username, password, endpoint string
	httpClient *http.Client
	Project *ProjectService
	Repositories *RepositoriesService
	Tags *TagsService
}

func NewClient(	username, password, endpoint string) (*Client) {
	c := &Client{httpClient: http.DefaultClient, username: username, password: password}
	if !strings.HasSuffix(endpoint, "/api") {
		endpoint = fmt.Sprintf("%v/api", endpoint)
	}
	c.endpoint = endpoint
	c.Project = &ProjectService{client: c}
	c.Repositories = &RepositoriesService{client: c}
	c.Tags = &TagsService{client: c}
	return c
}

func (c *Client) requestExtHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("User-Agent", "harbor-sync")
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	return Do(c.httpClient, req, v)
}

func Do(c *http.Client, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if v != nil {
		defer resp.Body.Close()
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			decoder := json.NewDecoder(resp.Body)
			// decoder.DisallowUnknownFields()
			err = decoder.Decode(v)
		}
	}
	return resp, err
}
