package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	setupRouter()
	server := newServer()
	server.ListenAndServe()
}

func newServer() *http.Server {
	def := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return def
}

//setup sets route
func setupRouter() {
	http.HandleFunc("localhost/healthz", healthzHandler)
	http.HandleFunc("/set/response", writeReponseHandler)
}

func healthzHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	accessLog(r, http.StatusOK)
}

func writeReponseHandler(rw http.ResponseWriter, r *http.Request) {
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
	accessLog(r, http.StatusOK)
}

//access records access log.
func accessLog(r *http.Request, statusCode int) {
	log.Printf("client_ip: %s, response_status_code: %+v", clientIP(r), statusCode)
}

func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
