package http

import (
	"log"
	"net/http"
	"time"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		log.Printf("[%s] %s  at: %s", req.Method, req.RequestURI, start)
	})
}
