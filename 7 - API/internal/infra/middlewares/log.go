package middlewares

import (
	"log"
	"net/http"
)

func CustomLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Custom logger: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
