package proxy

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

// Config : struct representation of a toml config
type Config struct {
	path        string
	destination int
	entrypoint  int
}

// LoadConfig  reads a toml file and returns a Config struct
func LoadConfig() *Config {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("could not load toml config")
	}

	proxyconf := Config{
		path:        config.Get("path").(string),
		destination: config.Get("destination").(int),
		entrypoint:  config.Get("entrypoint").(int),
	}
	return &proxyconf
}
