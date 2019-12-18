// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

type ContentTypeField struct {
	FieldType   string `json:"field_type"`
	ApiId       string `json:"api_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Localized   bool   `json:"is_localized"`
}

type ContentType struct {
	TypeId      string             `json:"type_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Fields      []ContentTypeField `json:"fields"`
}

type ContentTypeEntry struct {
	ApiId       string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EntryCount  int    `json:"entry_count"`
}

type ContentTypeCollection struct {
	Meta    CollectionMeta     `json:"meta"`
	Entries []ContentTypeEntry `json:"entries"`
}

// Fetches a collection of content types. Can use nil as filter if no filtering is needed
func (api *DeliveryApi) GetContentTypeCollection(filter *Filter) (*ContentTypeCollection, error) {
	qs := filter.createQueryString()
	url := fmt.Sprintf("/types?%s", qs)

	coll := ContentTypeCollection{}
	err := api.seamsClient.Fetch("GET", url, nil, &coll)
	if err != nil {
		return nil, err
	}

	return &coll, nil
}

// Fetches the given content type details
func (api *DeliveryApi) GetContentType(contentType string, filter *Filter) (*ContentType, error) {
	qs := filter.createQueryString()
	url := fmt.Sprintf("/type/%s?%s", contentType, qs)

	ct := ContentType{}
	err := api.seamsClient.Fetch("GET", url, nil, &ct)
	if err != nil {
		return nil, err
	}

	return &ct, nil
}
