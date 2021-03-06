package proxy

import (
	"log"

	"github.com/pelletier/go-toml"
)

// Config : struct representation of a toml config
type Config struct {
	Path        string
	Destination string
	Entrypoint  string
}

// NewConfig  Returns a pointer to a new config struct
func NewConfig() *Config {
	return &Config{}
}

// LoadConfig reads config.toml and returns a struct of type Config
func LoadConfig() *Config {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Panicf("could not load toml file %s", err.Error())
	}
	path := config.Get("proxy.path").(string)
	destination := config.Get("proxy.destination").(string)
	entrypoint := config.Get("proxy.entrypoint").(string)
	// todo: validate data

	parsedtoml := Config{
		Path:        path,
		Destination: destination,
		Entrypoint:  entrypoint,
	}
	return &parsedtoml

}
