// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

// Package client implements a generic API client for Seams-CMS
package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const UserAgent = "Go-seams-cms-sdk/v1.0"

type Client struct {
	Workspace string
	ApiKey    string
	BaseUrl   string
	Http      *http.Client
}

type Configuration struct {
	Workspace string
	ApiKey    string
	BaseUrl   string
}

func NewWithConfig(config Configuration) *Client {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpClient := &http.Client{Transport: transport}

	return &Client{
		ApiKey:    config.ApiKey,
		Workspace: config.Workspace,
		BaseUrl:   config.BaseUrl,
		Http:      httpClient,
	}
}

func (c *Client) Fetch(method string, url string, body io.Reader, v interface{}) error {
	req, err := c.createRequest(method, url, body)
	if err != nil {
		return err
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (c *Client) createRequest(method string, url string, body io.Reader) (*http.Request, error) {
	if url[0] != '/' {
		url = "/" + url
	}
	url = c.BaseUrl + "/workspace/" + c.Workspace + url

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	return req, nil
}
