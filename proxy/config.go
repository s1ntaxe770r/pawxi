package proxy

// struct representation of toml config
type ProxyConfig struct {
	path        string
	destination int
	entrypoint  int
}
