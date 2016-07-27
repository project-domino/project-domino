package config

// Database is a type for database configuration settings.
type Database struct {
	Debug bool   `toml:"debug"`
	Type  string `toml:"type"`
	URL   string `toml:"url"`
}

// DefaultDatabase is the default database configuration.
var DefaultDatabase = Database{
	Debug: false,
	Type:  "postgres",
	URL:   "dbname=domino host=domino sslmode=disable user=domino",
}
