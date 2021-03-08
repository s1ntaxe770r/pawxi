package tls

import "golang.org/x/crypto/acme/autocert"

// NewCertManager returns a new certmanager
func NewCertManager() *autocert.Manager {
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	return &certManager

}
