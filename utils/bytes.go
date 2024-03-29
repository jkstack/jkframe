package utils

import (
	"math"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"gopkg.in/yaml.v3"
)

// Bytes custom bytes struct
type Bytes uint64

// MarshalJSON marshal bytes by json
func (data Bytes) MarshalJSON() ([]byte, error) {
	return []byte(`"` + humanize.IBytes(uint64(data)) + `"`), nil
}

// UnmarshalJSON unmarshal bytes by json
func (data *Bytes) UnmarshalJSON(value []byte) error {
	str := strings.TrimPrefix(string(value), `"`)
	str = strings.TrimSuffix(str, `"`)
	n, err := humanize.ParseBytes(str)
	if err != nil {
		return err
	}
	*data = Bytes(n)
	return nil
}

// MarshalYAML marshal bytes by yaml
func (data Bytes) MarshalYAML() (interface{}, error) {
	return humanize.IBytes(uint64(data)), nil
}

// UnmarshalYAML unmarshal bytes by yaml
func (data *Bytes) UnmarshalYAML(value *yaml.Node) error {
	n, err := humanize.ParseBytes(value.Value)
	if err != nil {
		return err
	}
	*data = Bytes(n)
	return nil
}

// MarshalKV marshal bytes by kv
func (data *Bytes) MarshalKV() (string, error) {
	return humanize.IBytes(uint64(*data)), nil
}

// UnmarshalKV unmarshal bytes by kv
func (data *Bytes) UnmarshalKV(value string) error {
	n, err := humanize.ParseBytes(value)
	if err != nil {
		return err
	}
	*data = Bytes(n)
	return nil
}

// Bytes get bytes data
func (data *Bytes) Bytes() uint64 {
	return uint64(*data)
}

// String format to string
func (data Bytes) String() string {
	return humanize.IBytes(uint64(data))
}

// Number get number
func (data Bytes) Number() uint64 {
	str := humanize.IBytes(uint64(data))
	tmp := strings.SplitN(str, " ", 2)
	if len(tmp) == 2 {
		n, _ := strconv.ParseFloat(tmp[0], 64)
		return uint64(math.Ceil(n))
	}
	return 0
}

// Unit get unit
func (data Bytes) Unit() string {
	str := humanize.IBytes(uint64(data))
	tmp := strings.SplitN(str, " ", 2)
	if len(tmp) == 2 {
		return tmp[1]
	}
	return ""
}
