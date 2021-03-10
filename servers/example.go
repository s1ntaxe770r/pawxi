package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	sm := http.NewServeMux()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request recieved from %s", r.Host)
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}

		fmt.Println(string(dump))
		w.Write([]byte(`{"ping":"pong"}`))

	})

	log.Println("listening on 6000")
	log.Fatal(http.ListenAndServe(":6000", sm))

}
