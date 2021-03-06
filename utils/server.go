package utils

import (
	"net/http"
	"time"
)

// NewServer returns a new http server with custom timeouts
func NewServer(servemux *http.ServeMux, port string) *http.Server {
	// enforce proper timeouts on servver
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      servemux,
		Addr:         port,
	}

	return server
}
