package config

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	"github.com/naoina/toml"
)

// LoadConfig loads from one or more configuration files.
func LoadConfig(config interface{}, files []string) error {
	if len(files) == 0 {
		log.Println("WARNING: Using default config")
	}
	for _, file := range files {
		if err := loadConfigFile(config, file); err != nil {
			return err
		}
	}
	log.Printf("Loaded config %s", pp.Sprint(config))
	return nil
}

func loadConfigFile(config interface{}, file string) error {
	log.Printf("Loading config from file %s...", file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return toml.Unmarshal(b, config)
}
