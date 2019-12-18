// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFilter_withoutNil(t *testing.T) {
	sort := "foo"
	query := "bar"
	f := NewFilter(5, 10, &sort, &query)

	assert.Equal(t, 5, f.Offset)
	assert.Equal(t, 10, f.Limit)
	assert.Equal(t, "foo", *f.Sort)
	assert.Equal(t, "bar", *f.Query)
}

func TestNewFilter_withNil(t *testing.T) {
	f := NewFilter(5, 10, nil, nil)

	assert.Equal(t, 5, f.Offset)
	assert.Equal(t, 10, f.Limit)
	assert.Nil(t, f.Sort)
	assert.Nil(t, f.Query)
}

func TestNewFilter(t *testing.T) {
	type providerData struct {
		Offset	int
		Limit	int
		Sort	string
		Query	string
		Result	string
	}

	provider := []providerData{
		{5, 10, "", "", "offset=5&limit=10"},
		{0, 100, "", "", "offset=0&limit=100"},
		{0, 10, "foo=asc", "", "offset=0&limit=10&sort=foo%3Dasc"},
		{0, 10, "foo=asc", "c.f.d eq \"foo\"", "offset=0&limit=10&sort=foo%3Dasc&query=c.f.d+eq+%22foo%22"},
		{0, 10, "foo=asc", "c.f.d eq \"foo\" and c.d.a contains foobar", "offset=0&limit=10&sort=foo%3Dasc&query=c.f.d+eq+%22foo%22+and+c.d.a+contains+foobar"},
	}

	for _, data := range(provider) {
		var sort *string = nil
		if data.Sort != "" {
			sort = &data.Sort
		}
		var query *string = nil
		if data.Query != "" {
			query = &data.Query
		}

		f := NewFilter(data.Offset, data.Limit, sort, query)
		assert.Equal(t, data.Result, f.createQueryString())
	}
}
