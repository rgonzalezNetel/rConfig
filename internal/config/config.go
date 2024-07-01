package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Folders struct {
		Names []string `toml:"names"`
	} `toml:"folders"`

	Files struct {
		Paths []string `toml:"paths"`
	} `toml:"files"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return config, err
	}
	return config, nil
}
