// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

// Package change implements the Seams-CMS change API
package change

import "github.com/seams-cms/go-seams-cms-sdk/client"

const (
	changeApi      = "https://change.seams-api.com"
	noCdnChangeApi = "https://change-nocdn.seams-api.com"
)

type ChangeApi struct {
	seamsClient *client.Client
	baseUrl 	string
}

type Configuration struct {
	Workspace string
	ApiKey    string
	UseCdn    bool
	BaseUrl   string
}

func NewClient(apiKey string, workspace string) *ChangeApi {
	config := Configuration{
		ApiKey:    apiKey,
		Workspace: workspace,
		UseCdn:    true,
	}

	return NewClientWithConfig(&config)
}

func NewClientWithConfig(config *Configuration) *ChangeApi {
	baseUrl := changeApi
	if config.UseCdn == false {
		baseUrl = noCdnChangeApi
	}

	// Override base URL if needed
	if config.BaseUrl != "" {
		baseUrl = config.BaseUrl
	}

	apiClient := client.NewWithConfig(
		client.Configuration{
			ApiKey:    config.ApiKey,
			Workspace: config.Workspace,
			BaseUrl:   baseUrl,
		},
	)

	return &ChangeApi{
		seamsClient: apiClient,
	}
}
