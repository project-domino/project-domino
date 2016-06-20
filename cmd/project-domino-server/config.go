package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	// Set variable defaults.
	viper.SetDefault("assets.dev", false)
	viper.SetDefault("assets.path", "assets.zip")
	viper.SetDefault("database.debug", false)
	viper.SetDefault("database.type", "sqlite3")
	viper.SetDefault("database.url", "domino.db")
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
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.BindEnv("http.port", "PORT")

	// Read config or die.
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
