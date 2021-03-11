package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/NYTimes/gziphandler"
	"github.com/fsnotify/fsnotify"
	"github.com/s1ntaxe770r/pawxi/proxy"
	"github.com/s1ntaxe770r/pawxi/tls"
	"github.com/s1ntaxe770r/pawxi/utils"
	"github.com/spf13/viper"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	r      = http.NewServeMux()
	routes []proxy.Route
)

func main() {

	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetConfigName("pawxi")
	readerr := viper.ReadInConfig()
	port := fmt.Sprintf(":%s", viper.GetString("proxy.binds"))
	server := utils.NewServer(r, port)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

	})
	if readerr != nil {
		panic(fmt.Errorf("fatal error config file: %s", readerr))
	}
	err := viper.UnmarshalKey("proxy.routes", &routes)
	if err != nil {
		log.Fatalf("unable to unmarshal into strcut REASON : %v", err.Error())
	}
	usegzip := viper.GetBool("proxy.usegzip")
	for _, route := range routes {
		// parse each url
		dest, err := url.Parse(route.Destination)
		if err != nil {
			log.Fatalf("could not parse url %s", route.Destination)
		}
		// indiviually proxy each request
		proxy := proxy.Tunnel(route, dest)
		if usegzip != true {
			go r.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
				proxy.ServeHTTP(w, r)
			})
		} else {
			zipped := gziphandler.GzipHandler(proxy)
			go r.Handle(route.Path, zipped)
		}
	}
	tls.SetupDevCerts()
	vendorPath := "./certs"
	cert := filepath.Join(vendorPath, "devcerts", "cert.pem")
	key := filepath.Join(vendorPath, "devcerts", "key.pem")
	go func() {
		if err := http.ListenAndServe(port, http.HandlerFunc(utils.RedirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	server.ListenAndServeTLS(cert, key)

}
