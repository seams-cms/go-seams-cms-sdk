// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

// Package delivery implements the Seams-CMS delivery API
package delivery

import (
	client "github.com/seams-cms/go-seams-cms-sdk/client"
)

const (
	deliveryApi      = "https://delivery.seams-api.com"
	noCdnDeliveryApi = "https://delivery-nocdn.seams-api.com"
)

type DeliveryApi struct {
	seamsClient *client.Client
}

type Configuration struct {
	Workspace string
	ApiKey    string
	UseCdn    bool
	BaseUrl   string
}

// Creates a new delivery client based on workspace and api key
func NewClient(workspace string, apiKey string) *DeliveryApi {
	config := Configuration{
		ApiKey:    apiKey,
		Workspace: workspace,
		UseCdn:    true,
	}

	return NewClientWithConfig(&config)
}

// Creates a new delivery client based on custom configuration
func NewClientWithConfig(config *Configuration) *DeliveryApi {
	baseUrl := deliveryApi
	if config.UseCdn == false {
		baseUrl = noCdnDeliveryApi
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

	return &DeliveryApi{
		seamsClient: apiClient,
	}
}
