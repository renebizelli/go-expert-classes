package main

import "net/http"

type server struct{}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.Handle("/server", server{})
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok handler"))
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok ServeHTTP"))
}
