// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

import (
	"fmt"
	"net/url"
	"strings"
)

type Filter struct {
	Offset int
	Limit  int
	Sort   *string
	Query  *string
}

// Creates a new filter with given values
func NewFilter(offset int, limit int, sort *string, query *string) *Filter {
	return &Filter{
		Offset: offset,
		Limit:  limit,
		Sort:   sort,
		Query:  query,
	}
}

// Returns a query string snippet that can be used for calling the API
func (f *Filter) createQueryString() string {
	items := []string{
		fmt.Sprintf("offset=%d", f.Offset),
		fmt.Sprintf("limit=%d", f.Limit),
	}

	if f.Sort != nil {
		items = append(items, fmt.Sprintf("sort=%s", url.QueryEscape(*f.Sort)))
	}
	if f.Query != nil {
		items = append(items, fmt.Sprintf("query=%s", url.QueryEscape(*f.Query)))
	}

	return strings.Join(items, "&")
}
