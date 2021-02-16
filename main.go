package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/fatih/color"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	target, err := url.Parse("http://localhost:6000")
	handle(err)

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", handler(proxy))

	serverport, isSet := os.LookupEnv("PROXY_PORT")
	if isSet == true {
		green := color.New(color.FgGreen).PrintFunc()
		green("proxying on %s", serverport)
		log.Fatal(http.ListenAndServe(":"+serverport, nil))
	}
	color.Green("proxying on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}
