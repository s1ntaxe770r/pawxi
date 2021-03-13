package proxy

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// Tunnel : forward traffic to other ports
func Tunnel(route Route) http.Handler {
	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		dest, err := url.Parse(route.Destination)
		if err != nil {
			panic("unable to parse url")
		}
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", route.Destination)
		req.Host = dest.Host
		req.URL.Host = dest.Host
		req.URL.Scheme = dest.Scheme
		req.URL.Path = dest.Path
	},
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
		},
	}
	return proxy
}
