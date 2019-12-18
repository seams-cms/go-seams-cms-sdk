// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

import (
	"fmt"
	"time"
)

type AssetEntry struct {
	Asset struct {
		Workspace     string `json:"workspace"`
		Link          string `json:"link"`
		ThumbnailLink string `json:"thumbnail_link"`
		Size          int    `json:"size"`
		Path          string `json:"path"`
		Title         string `json:"title"`
		MimeType      string `json:"mimetype"`
	} `json:"asset"`
	Meta struct {
		CreatedAt time.Time `json:"created_at"`
		CreatedBy string    `json:"created_by"`
		UpdatedAt time.Time `json:"updated_at"`
		UpdatedBy string    `json:"updated_by"`
	} `json:"meta"`
}

type AssetCollection struct {
	Meta    CollectionMeta `json:"meta"`
	Entries []AssetEntry   `json:"entries"`
}

// Fetches assets. Can use nil as filter if no filtering is needed
func (api *DeliveryApi) GetAssetCollection(filter *Filter) (*AssetCollection, error) {
	qs := filter.createQueryString()
	url := fmt.Sprintf("/assets?%s", qs)

	coll := AssetCollection{}
	err := api.seamsClient.Fetch("GET", url, nil, &coll)
	if err != nil {
		return nil, err
	}

	return &coll, nil
}
