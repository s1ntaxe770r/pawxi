package proxy

// Config : struct representation of a toml config
type Config struct {
	Domain  string
	UseGzip string
}

// Route maps a path to a backend or destination. Eg: /home => http:localhost:6000/
type Route struct {
	Path, Destination string
}

// ProxyRoutes a slice of Route
var ProxyRoutes []Route

// NewConfig  Returns a pointer to a new config struct
