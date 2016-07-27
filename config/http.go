package config

// HTTP is a type for HTTP server settings.
type HTTP struct {
	Debug bool   `toml:"debug"`
	Dev   bool   `toml:"dev"`
	Port  uint16 `toml:"port"`
}

// DefaultHTTP is the default HTTP server configuration.
var DefaultHTTP = HTTP{
	Debug: false,
	Dev:   false,
	Port:  80,
}
