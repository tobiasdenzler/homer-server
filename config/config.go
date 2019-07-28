package config

import (
	"log"

	"github.com/jinzhu/configor"
)

// Config is based on config/config.yaml
var Config = struct {
	AppName string

	Server struct {
		Address    string
		LoginToken string
	}
}{}

func init() {
	// Load configuration file
	configor.Load(&Config, "config/config.yaml")
	log.Println("Config file loaded")
}
