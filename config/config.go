package config

import (
	"io/ioutil"

	"github.com/naoina/toml"
)

// LoadConfig loads from one or more configuration files.
func LoadConfig(config interface{}, files []string) error {
	for _, file := range files {
		if err := loadConfigFile(config, file); err != nil {
			return err
		}
	}
	return nil
}

func loadConfigFile(config interface{}, file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return toml.Unmarshal(b, config)
}
