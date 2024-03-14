package middleware

import (
	"log"
	"net/http"
)

func WithLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path:%s process start...\n", r.URL.Path)
		defer func() {
			log.Printf("path:%s process end...\n", r.URL.Path)
		}()
		handler.ServeHTTP(w, r)
	})
}
