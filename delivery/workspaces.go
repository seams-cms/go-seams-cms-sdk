// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

type WorkspaceEntry struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsArchived   bool   `json:"is_archived"`
	Organisation string `json:"organisation"`
	Count        struct {
		Types   int `json:"types"`
		Entries int `json:"entries"`
		Assets  int `json:"assets"`
	} `json:"count"`
	Locales []struct {
		Name   string `json:""`
		Locale string `json:""`
	} `json:"locales"`
}

type WorkspaceCollection struct {
	Meta    CollectionMeta   `json:"meta"`
	Entries []WorkspaceEntry `json:"entries"`
}

// Fetches a collection of workspaces (currently only one workspace is supported)
func (api *DeliveryApi) GetWorkspaceCollection() (*WorkspaceCollection, error) {
	coll := WorkspaceCollection{}
	err := api.seamsClient.Fetch("GET", "/", nil, &coll)
	if err != nil {
		return nil, err
	}

	return &coll, nil
}
