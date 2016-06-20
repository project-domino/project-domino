package main

import "github.com/spf13/viper"

type configAssets struct {
	Dev  bool   `uniconf:"dev"`
	Path string `uniconf:"path"`
}
type configDB struct {
	Addr string `uniconf:"url"`
	Type string `uniconf:"type"`

	Debug bool `uniconf:"debug"`
}

func init() {
	viper.SetDefault("assets.dev", false)
	viper.SetDefault("assets.path", "assets.zip")
	viper.SetDefault("db.addr", "domino.db")
	viper.SetDefault("db.type", "sqlite3")
	viper.SetDefault("db.debug", false)

	viper.SetConfigName("project-domino")
	viper.AddConfigPath("/etc/project-domino/")
	viper.AddConfigPath("$HOME/.project-domino")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
