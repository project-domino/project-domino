package config

// SMTP is a type for SMTP server settings.
type SMTP struct {
	Address string `toml:"address"`
	Host    string `toml:"host"`

	Identity string `toml:"identity"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

// DefaultSMTP is the default SMTP server configuration.
var DefaultSMTP = SMTP{
	Address:  "do-not-reply@project-domino.com",
	Host:     "smtp:25",
	Identity: "",
	Username: "domino",
	Password: "",
}
