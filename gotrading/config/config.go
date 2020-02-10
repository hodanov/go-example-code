package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// List reads config from config.ini.
type List struct {
	APIKey    string
	APISecret string
}

// Config ...
var Config List

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = List{
		APIKey:    cfg.Section("bitflyer").Key("api_key").String(),
		APISecret: cfg.Section("bitflyer").Key("api_secret").String(),
	}
}
