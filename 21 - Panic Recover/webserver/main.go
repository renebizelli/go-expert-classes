package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entrou no middleware")

		defer func() {

			if err := recover(); err != nil {
				log.Printf("Erro: %v", err)
				debug.PrintStack() // Imprime o stack trace
				http.Error(w, "Ocorreu um erro", http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)

	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola mundo"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Error")
	})

	if h := http.ListenAndServe(":8080", recoverMiddleware(mux)); h != nil {
		log.Fatalf("Could not listen on %s: %v\n", ":8080", h)
	}

}
