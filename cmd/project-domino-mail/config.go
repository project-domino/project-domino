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
	viper.SetDefault("database", map[string]interface{}{
		"debug": false,
		"type":  "postgres",
		"url":   "dbname=domino sslmode=disable",
	})
	viper.SetDefault("smtp", map[string]interface{}{
		"address":  "localhost:25",
		"identity": "",
		"username": "domino",
		"password": "",
		"host":     "localhost:25",
	})

	// Setup configuration files.
	viper.SetConfigName("project-domino-mail")
	viper.AddConfigPath("/etc/project-domino-mail/")
	viper.AddConfigPath("$HOME/.project-domino-mail")
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
