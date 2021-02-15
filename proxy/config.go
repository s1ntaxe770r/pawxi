package proxy

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

// struct representation of toml config
type ProxyConfig struct {
	path        string
	destination int
	entrypoint  int
}

func LoadConfig() *ProxyConfig {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("could not load toml config")
	}

	proxyconf := ProxyConfig{path: config.Get("path").(string), destination: config.Get("destination").(string)}

}
