package config

import (
	"github.com/jinzhu/configor"
)

// Config is based on config/config.yaml
var Config = struct {
	Server struct {
		Address    string
		LoginToken string
	}
	Log struct {
		Level string
	}
}{}

func init() {
	// Load configuration file
	// Path must be relative to the base directory
	LoadConfigFile("config/config.yaml")
}

// LoadConfigFile loads the config file from a specified location.
// Can also be used for running unit tests.
func LoadConfigFile(path string) error {
	return configor.Load(&Config, path)
}
