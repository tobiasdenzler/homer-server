package dss

import "testing"

func TestCreateParamString(t *testing.T) {
	params := map[string]string{"a": "x", "b": "y"}
	result := createParamString(params)
	if result != "a=x&b=y" && result != "b=y&a=x" {
		t.Error("Expected a=x&b=y or b=y&a=x, got ", result)
	}
}
