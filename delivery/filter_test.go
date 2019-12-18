// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

import "testing"

func TestNewFilter_withoutNil(t *testing.T) {
	sort := "foo"
	query := "bar"
	f := NewFilter(5, 10, &sort, &query)

	if (f.Offset != 5) {
		t.Errorf("offset %d, want %d", f.Offset, 5)
	}

	if (f.Limit != 10) {
		t.Errorf("limit %d, want %d", f.Limit, 10)
	}

	if (*f.Sort != "foo") {
		t.Errorf("sort %s, want %s", *f.Sort, "foo")
	}

	if (*f.Query != "bar") {
		t.Errorf("query %s, want %s", *f.Query, "bar")
	}
}

func TestNewFilter_withNil(t *testing.T) {
	f := NewFilter(5, 10, nil, nil)

	if (f.Offset != 5) {
		t.Errorf("offset %d, want %d", f.Offset, 5)
	}

	if (f.Limit != 10) {
		t.Errorf("limit %d, want %d", f.Limit, 10)
	}

	if (f.Sort != nil) {
		t.Errorf("sort needs to be nil")
	}

	if (f.Query != nil) {
		t.Errorf("query needs to be nil")
	}
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

		qs := f.createQueryString()
		if (qs != data.Result) {
			t.Errorf("query string %s, want %s", qs, data.Result)
		}
	}
}
