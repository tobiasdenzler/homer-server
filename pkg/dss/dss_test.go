package dss

import (
	"log"
	"testing"

	"github.com/tobiasdenzler/homer-server/config"
)

func init() {
	config.LoadConfigFile("../../config/config.yaml")
	log.Println("Loaded config ->", config.Config)
}

func TestCreateParamString(t *testing.T) {
	params := map[string]string{"a": "x", "b": "y"}
	result := createParamString(params)
	if result != "a=x&b=y" && result != "b=y&a=x" {
		t.Errorf("Query string is incorrect, got: %s, want: %s", result, "a=x&b=y or b=y&a=x")
	}
}

func TestLogin(t *testing.T) {
	err := login()
	if err != nil {
		t.Error("No error expected, got ", err)
	}
}
