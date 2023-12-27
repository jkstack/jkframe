package utils

import (
	"fmt"
)

// Float64P2 format yaml, json, kv string to 2 precision
type Float64P2 float64

// MarshalJSON marshal float64 by json
func (data Float64P2) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", data)), nil
}

// MarshalYAML marshal float64 by yaml
func (data Float64P2) MarshalYAML() (interface{}, error) {
	return fmt.Sprintf("%.2f", data), nil
}

// MarshalKV marshal float64 by kv
func (data Float64P2) MarshalKV() (string, error) {
	return fmt.Sprintf("%.2f", data), nil
}

// Float get float64 data
func (data Float64P2) Float() float64 {
	return float64(data)
}

// String get string data
func (data Float64P2) String() string {
	return fmt.Sprintf("%.2f", data)
}
