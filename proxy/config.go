package proxy

import "github.com/pelletier/go-toml"

// struct representation of toml config
type ProxyConfig struct {
	path        string
	destination int
	entrypoint  int
}
 
func LoadConfig() *ProxyConfig {
	config, err := go-toml.Load("config.toml")
	if err != nil {
		fmt.Println("could not load toml config")
	}

}
