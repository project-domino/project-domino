package main

import "github.com/remexre/uniconf"

type configAssets struct {
	Dev  bool   `uniconf:"dev"`
	Path string `uniconf:"path"`
}
type configDB struct {
	Addr string `uniconf:"url"`
	Type string `uniconf:"type"`

	Debug bool `uniconf:"debug"`
}

var config = struct {
	Assets configAssets
	DB     configDB
}{
	Assets: configAssets{
		Dev:  false,
		Path: "assets.zip",
	},
	DB: configDB{
		Addr:  "domino.db",
		Type:  "sqlite3",
		Debug: false,
	},
}

func init() {
	uniconf.MustLoad(&config)
}
