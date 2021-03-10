package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Tunnel : forward traffic to other ports
func Tunnel(route Route, dest *url.URL) http.Handler {
	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", dest.Host)
		req.Host = dest.Host
		req.URL.Host = dest.Host
		req.URL.Scheme = dest.Scheme

	}}

	return proxy
}
