// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

import (
	"encoding/json"
	"fmt"
	"time"
)

var convert = Converter{}

type ContentEntry struct {
	Meta struct {
		EntryId    string    `json:"entry_id"`
		RevisionId string    `json:"revision_id"`
		TypeId     string    `json:"type_id"`
		CreatedAt  time.Time `json:"created_at"`
		CreatedBy  string    `json:"created_by"`
		UpdatedAt  time.Time `json:"updated_at"`
		UpdatedBy  string    `json:"updated_by"`
	} `json:"meta"`
	Content map[string]Content `json:"content"`
}

type ContentCollection struct {
	Meta    CollectionMeta `json:"meta"`
	Entries []ContentEntry `json:"entries"`
}

type LocaleType map[string]interface{}

type Content struct {
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
	Locales LocaleType  `json:"locales"`
}

// PHP has a way to display empty (dict) arrays as []. This confuses Go so we need to manually check for this.
func (l *LocaleType) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	switch t := v.(type) {
	// Found (an empty) [], convert to empty LocaleType
	case []interface{}:
		*l = LocaleType{}
	case map[string]interface{}:
		*l = LocaleType(t)
	default:
		return fmt.Errorf("unknown locale type found: %s", t)
	}

	return nil
}

// Returns true when the content is localized
func (c *Content) IsLocalized() bool {
	return len(c.Locales) > 0
}

// Returns true when the given data is actual content
func (c *Content) IsContent() bool {
	return c.Type == "content"
}

// Returns true when the given data is a reference to another content entry
func (c *Content) IsReference() bool {
	return c.Type == "reference"
}

// Returns sliced data or error when given content could not be sliced
func (c *Content) GetSlice() ([]interface{}, error) {
	a, err := convert.ToSlice(c.Value)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Returns string data or error when given content could not be converted to a string
func (c *Content) GetString() (string, error) {
	s, err := convert.ToString(c.Value)
	if err != nil {
		return "", err
	}

	return s, nil
}

// Returns int64 data or error when given content could not be converted to an int64
func (c *Content) GetInt() (int64, error) {
	i, err := convert.ToInt64(c.Value)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// Returns bool data or error when given content could not be converted to a bool
func (c *Content) GetBool() (bool, error) {
	b, err := convert.ToBool(c.Value)
	if err != nil {
		return false, err
	}

	return b, nil
}

// Returns slice of references or error when content is not a reference slice
func (c *Content) GetReferences() ([]ContentEntry, error) {
	var entries []ContentEntry
	buf, err := json.Marshal(c.Value)
	err = json.Unmarshal(buf, &entries)
	if err != nil {
		return entries, err
	}

	return entries, nil
}

// Returns slice for given localized content
func (c *Content) GetLocaleSlice(locale string) ([]interface{}, error) {
	if _, ok := c.Locales[locale]; !ok {
		return nil, fmt.Errorf("unknown locale found: %s", locale)
	}

	a, err := convert.ToSlice(c.Locales[locale])
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Returns string for given localized content
func (c *Content) GetLocaleString(locale string) (string, error) {
	if _, ok := c.Locales[locale]; !ok {
		return "", fmt.Errorf("unknown locale found: %s", locale)
	}

	s, err := convert.ToString(c.Locales[locale])
	if err != nil {
		return "", err
	}

	return s, nil
}

// Returns int64 for given localized content
func (c *Content) GetLocaleInt(locale string) (int64, error) {
	if _, ok := c.Locales[locale]; !ok {
		return 0, fmt.Errorf("unknown locale found: %s", locale)
	}

	i, err := convert.ToInt64(c.Locales[locale])
	if err != nil {
		return 0, err
	}

	return i, nil
}

// Returns bool for given localized content
func (c *Content) GetLocaleBool(locale string) (bool, error) {
	if _, ok := c.Locales[locale]; !ok {
		return false, fmt.Errorf("unknown locale found: %s", locale)
	}

	b, err := convert.ToBool(c.Locales[locale])
	if err != nil {
		return false, err
	}

	return b, nil
}

// Returns reference slice for given localized content
func (c *Content) GetLocaleReferences(locale string) ([]ContentEntry, error) {
	if _, ok := c.Locales[locale]; !ok {
		return nil, fmt.Errorf("unknown locale found: %s", locale)
	}

	var entries []ContentEntry
	buf, err := json.Marshal(c.Locales[locale])
	err = json.Unmarshal(buf, &entries)
	if err != nil {
		return entries, err
	}

	return entries, nil
}

// Fetches a collection of content entries based on content type. Can use nil as filter if no filtering is needed
func (api *DeliveryApi) GetContentCollection(contentType string, filter *Filter) (*ContentCollection, error) {
	qs := filter.createQueryString()
	url := fmt.Sprintf("/type/%s/entries?%s", contentType, qs)

	coll := ContentCollection{}
	err := api.seamsClient.Fetch("GET", url, nil, &coll)
	if err != nil {
		return nil, err
	}

	return &coll, nil
}

// Fetches content entry element based on entry ID
func (api *DeliveryApi) GetContent(entryId string) (*ContentEntry, error) {
	entry := ContentEntry{}
	return &entry, nil
}
