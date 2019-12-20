// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	api := NewClient("space", "api-key")

	assert.Contains(t, api.seamsClient.BaseUrl, "https://delivery.seams-api.com")
}

func TestNewClientWithConfig(t *testing.T) {
	config := Configuration{
		"space",
		"api-key",
		false,
		"",
	}
	api := NewClientWithConfig(&config)

	assert.Contains(t, api.seamsClient.BaseUrl, "https://delivery-nocdn.seams-api.com")
}

func TestNewClientWithConfig_baseurl(t *testing.T) {
	config := Configuration{
		"space",
		"api-key",
		false,
		"http://base.url",
	}
	api := NewClientWithConfig(&config)

	assert.Contains(t, api.seamsClient.BaseUrl, "http://base.url")
}
