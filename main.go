package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/NYTimes/gziphandler"
	"github.com/s1ntaxe770r/pawxi/proxy"
	"github.com/s1ntaxe770r/pawxi/utils"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := proxy.LoadConfig()
	port := config.Entrypoint
	usegzip := config.UseGzip
	server := utils.NewServer(nil, fmt.Sprintf(":%s", port))

	target, err := url.Parse(config.Destination)
	handle(err)

	proxy := httputil.NewSingleHostReverseProxy(target)
	if usegzip == "true" {
		zipped := gziphandler.GzipHandler(handler(proxy))
		http.Handle(config.Path, zipped)
		fmt.Printf("proxying on %s", port)
		server.ListenAndServe()
	}
	http.HandleFunc("/", handler(proxy))
	server.ListenAndServe()

}

func handler(p *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}
