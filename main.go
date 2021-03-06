package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/s1ntaxe770r/pawxi/proxy"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := proxy.LoadConfig()
	port := config.Entrypoint
	target, err := url.Parse("http://localhost:6000")
	handle(err)

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", handler(proxy))
	fmt.Printf("path = %s", config.Path)
	fmt.Printf("proxying on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}
