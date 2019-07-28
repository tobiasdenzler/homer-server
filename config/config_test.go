package config

import (
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	LoadConfigFile("config.template.yaml")
	if Config.Server.Address != "https://127.0.0.1:8080" {
		t.Errorf("Configuration value not loaded correctly, got: %s, want: %s", Config.Server.Address, "https://127.0.0.1:8080")
	}
}
