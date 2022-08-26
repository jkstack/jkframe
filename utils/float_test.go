package utils

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestFloatJson(t *testing.T) {
	var obj struct {
		Value Float64P2 `json:"value"`
	}
	obj.Value = Float64P2(1. / 3.)
	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, []byte(`{"value":0.33}`)) {
		t.Fatal("unexpected value")
	}
}
