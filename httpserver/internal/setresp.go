package internal

import (
	"net/http"
	"os"
)

// SetReponseHandler .
func SetReponseHandler(rw http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for k, values := range headers {
		for _, v := range values {
			rw.Header().Add(k, v)
		}
	}
	version := os.Getenv("VERSION")
	if version == "" {
		version = "NON-VERSION"
	}
	rw.Header().Add("VERSION", version)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("SetReponseHandler success"))
	AccessLog(r, http.StatusOK)
}
