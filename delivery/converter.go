// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package delivery

/* Conversion code inspired by https://github.com/icza/dyno/blob/master/dyno.go */

import "fmt"

type Converter struct {
}

// Converts given value to string
func (c *Converter) ToString(v interface{}) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("expected string value, got: %T", v)
	}

	return s, nil
}

// Converts given value to int64
func (c *Converter) ToInt64(v interface{}) (int64, error) {
	switch i := v.(type) {
	case int64:
		return i, nil
	case int:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case int16:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case uint:
		return int64(i), nil
	case uint64:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case uint16:
		return int64(i), nil
	case uint8:
		return int64(i), nil
	case float64:
		return int64(i), nil
	case float32:
		return int64(i), nil
	case string:
		var n int64
		_, err := fmt.Sscan(i, &n)
		return n, err
	case interface {
		Int64() (int64, error)
	}:
		return i.Int64()
	default:
		return 0, fmt.Errorf("expected some form of integer number, got: %T", v)
	}
}

// Converts given value to float64
func (c *Converter) ToFloat64(v interface{}) (float64, error) {
	switch f := v.(type) {
	case float64:
		return f, nil
	case float32:
		return float64(f), nil
	case int64:
		return float64(f), nil
	case int:
		return float64(f), nil
	case int32:
		return float64(f), nil
	case int16:
		return float64(f), nil
	case int8:
		return float64(f), nil
	case uint:
		return float64(f), nil
	case uint64:
		return float64(f), nil
	case uint32:
		return float64(f), nil
	case uint16:
		return float64(f), nil
	case uint8:
		return float64(f), nil
	case string:
		var n float64
		_, err := fmt.Sscan(f, &n)
		return n, err
	case interface {
		Float64() (float64, error)
	}:
		return f.Float64()
	default:
		return 0, fmt.Errorf("expected some form of floating point number, got: %T", v)
	}
}

// Converts given value to a boolean
func (c *Converter) ToBool(v interface{}) (bool, error) {
	switch f := v.(type) {
	case bool:
		return f, nil
	case int:
		return f != 0, nil
	case int64:
		return f != 0, nil
	case int32:
		return f != 0, nil
	case int16:
		return f != 0, nil
	case int8:
		return f != 0, nil
	case uint:
		return f != 0, nil
	case uint64:
		return f != 0, nil
	case uint32:
		return f != 0, nil
	case uint16:
		return f != 0, nil
	case uint8:
		return f != 0, nil
	case float64:
		return f != 0, nil
	case float32:
		return f != 0, nil
	case string:
		var n bool
		_, err := fmt.Sscan(f, &n)
		return n, err
	case interface {
		Float64() (float64, error)
	}:
		val, err := f.Float64()
		if err != nil {
			return false, err
		}
		return val != 0, err
	default:
		return false, fmt.Errorf("expected bool, got: %T", v)
	}
}

// Converts given value to int
func (c *Converter) ToInt(v interface{}) (int, error) {
	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("expected int value, got: %T", v)
	}

	return i, nil
}

// Converts given value to []interface{}
func (c *Converter) ToSlice(v interface{}) ([]interface{}, error) {
	s, ok := v.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected slice node, got: %T", v)
	}

	return s, nil
}

// Converts given value to a map[string]interface{}
func (c *Converter) ToMap(v interface{}) (map[string]interface{}, error) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected map with string keys node, got: %T", v)
	}

	return m, nil
}
