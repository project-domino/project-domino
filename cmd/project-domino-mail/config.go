package main

import (
	"flag"
	"log"

	"github.com/project-domino/project-domino/config"
)

// Config is the configuration for the server.
var Config ConfigType

// ConfigType is the type for the configuration for the server.
type ConfigType struct {
	Database config.Database `toml:"database"`
	SMTP     config.SMTP     `toml:"smtp"`
}

func init() {
	// Create default config object.
	Config = ConfigType{
		Database: config.DefaultDatabase,
		SMTP:     config.DefaultSMTP,
	}

	// Read config or die.
	if err := config.LoadConfig(Config, flag.Args()); err != nil {
		log.Fatal(err)
	}
}
