package internal

import "net/http"

// HealthHandler .
func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Health success"))
	AccessLog(r, http.StatusOK)
}
