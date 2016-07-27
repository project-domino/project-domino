package config

// Assets is a type for asset configuration settings.
type Assets struct {
	Dev  bool   `toml:"dev"`
	Path string `toml:"path"`
}

// DefaultAssets is the default database configuration.
var DefaultAssets = Assets{
	Dev:  false,
	Path: "assets.zip",
}
