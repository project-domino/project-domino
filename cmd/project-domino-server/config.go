package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	// Set variable defaults.
	viper.SetDefault("assets", map[string]interface{}{
		"dev":  false,
		"path": "assets.zip",
	})
	viper.SetDefault("database", map[string]interface{}{
		"debug": false,
		"type":  "postgres",
		"url":   "dbname=domino sslmode=disable",
	})
	viper.SetDefault("http", map[string]interface{}{
		"debug": false,
		"port":  80,
	})

	// Setup configuration files.
	viper.SetConfigName("project-domino")
	viper.AddConfigPath("/etc/project-domino/")
	viper.AddConfigPath("$HOME/.project-domino")
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(os.Args[0]))

	// Add environment variables.
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.BindEnv("http.port", "PORT")

	// Read config or die.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
