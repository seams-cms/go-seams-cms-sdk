// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package client

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestNewWithConfig(t *testing.T) {
	config := Configuration{
		"workspace",
		"api-key",
		"https://base.url/api",
	};
	c := NewWithConfig(config)

	assert.Equal(t, c.apiKey, "api-key")
	assert.Equal(t, c.workspace, "workspace")
	assert.Equal(t, c.baseUrl, "https://base.url/api")
}

func TestCreateRequest(t *testing.T) {
	config := Configuration{
		"myspace",
		"api-key",
		"https://base.url/api",
	};
	client := NewWithConfig(config)

	req, _ := client.createRequest("GET", "foobar", nil)
	assert.Equal(t, "https://base.url/api/workspace/myspace/foobar", req.URL.String())
}