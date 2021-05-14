package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/pawxi/metrics"
	"github.com/s1ntaxe770r/pawxi/proxy"
	"github.com/s1ntaxe770r/pawxi/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	router = mux.NewRouter()
	routes []proxy.Route
)

func main() {

	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetConfigName("pawxi")
	readerr := viper.ReadInConfig()
	port := fmt.Sprintf(":%s", viper.GetString("proxy.binds"))
	server := utils.NewServer(router, port)
	router.Use(metrics.Count_Request)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Warn("Config file changed:", e.Name)
		logrus.Warn("Restarting Proxy")
		s := spinner.New(spinner.CharSets[2], 100*time.Millisecond)
		s.Start()
		s.Color("yellow")
		// Semi Graceful shutdown ?
		time.Sleep(5 * time.Second)
		s.Stop()
		utils.Exec()
	})
	if readerr != nil {
		panic(fmt.Errorf("fatal error config file: %s", readerr))
	}
	err := viper.UnmarshalKey("proxy.routes", &routes)
	if err != nil {
		logrus.Error("unable to unmarshal into strcut REASON : %v", err.Error())
	}
	usegzip := viper.GetBool("proxy.usegzip")
	for _, route := range routes {
		// indiviually proxy each request
		proxy := proxy.Tunnel(route)
		if usegzip != true {
			go router.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
				proxy.ServeHTTP(w, r)

			})
		} else {
			zipped := gziphandler.GzipHandler(proxy)
			router.Handle(route.Path, zipped)
		}
	}
	utils.Vizualize(routes)
	logrus.Infof(color.HiGreenString("proxy started on %s"), port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}

}
