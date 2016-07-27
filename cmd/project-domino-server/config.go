package main

import (
	"log"
	"os"

	"github.com/project-domino/project-domino/config"
)

// Config is the configuration for the server.
var Config ConfigType

// ConfigType is the type of the configuration for the server.
type ConfigType struct {
	Assets   config.Assets   `toml:"assets"`
	Database config.Database `toml:"database"`
	HTTP     config.HTTP     `toml:"http"`
}

func init() {
	// Create default config object.
	Config = ConfigType{
		Assets:   config.DefaultAssets,
		Database: config.DefaultDatabase,
		HTTP:     config.DefaultHTTP,
	}

	// Read config or die.
	if err := config.LoadConfig(Config, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
