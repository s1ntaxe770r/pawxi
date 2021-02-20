package utils

import (
	"net/http"
	"time"
)

func NewServer(servemux *http.ServeMux) *http.Server {
	// enforce proper timeouts on servver
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      servemux,
	}

	return server
}
