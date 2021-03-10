package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/s1ntaxe770r/pawxi/proxy"
	"github.com/s1ntaxe770r/pawxi/utils"

	"github.com/spf13/viper"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := http.NewServeMux()
	server := utils.NewServer(r, ":80")
	var routes []proxy.Route
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetConfigName("pawxi")
	readerr := viper.ReadInConfig()
	if readerr != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", readerr))
	}

	err := viper.UnmarshalKey("proxy.routes", &routes)
	if err != nil {
		log.Fatalf("unable to unmarshal into strcut REASON : %v", err.Error())
	}

	for _, route := range routes {
		// parse each url
		dest, err := url.Parse(route.Destination)
		if err != nil {
			log.Fatalf("could not parse url %s", route.Destination)
		}
		// indiviually proxy each request
		proxy := proxy.Tunnel(route, dest)
		r.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	}

	log.Println("PROXYING ON PORT 80")
	server.ListenAndServe()

}
