package metrics

import (
	"expvar"
	"net/http"
	"time"

	"github.com/paulbellamy/ratecounter"
)

var (
	total_requests     = expvar.NewInt("total_requests")
	counter            = ratecounter.NewRateCounter(1 * time.Minute)
	request_per_minute = expvar.NewInt("hits_per_minute")
)

func Count_Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		counter.Incr(1)
		request_per_minute.Add(counter.Rate())
		total_requests.Set(1)
	})
}
