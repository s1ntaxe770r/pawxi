package utils

import (
	"net/http"
)

// RedirectTLS route incoming http trrafic => https
func RedirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https:"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}
