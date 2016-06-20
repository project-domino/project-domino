package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	// Set variable defaults.
	viper.SetDefault("assets.dev", false)
	viper.SetDefault("assets.path", "assets.zip")
	viper.SetDefault("db.addr", "domino.db")
	viper.SetDefault("db.type", "sqlite3")
	viper.SetDefault("db.debug", false)
	viper.SetDefault("http.debug", false)
	viper.SetDefault("http.port", 80)

	// Setup configuration files.
	viper.SetConfigName("project-domino")
	viper.AddConfigPath("/etc/project-domino/")
	viper.AddConfigPath("$HOME/.project-domino")
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(os.Args[0]))

	// Add environment variables.
	viper.AutomaticEnv()
	viper.BindEnv("http.port", "PORT")

	// Read config or die.
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
