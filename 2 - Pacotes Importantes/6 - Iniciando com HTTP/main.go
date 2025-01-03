package main

import "net/http"

func main() {

	http.HandleFunc("/cep", CepHandler)

	http.ListenAndServe(":8080", nil)

}

func CepHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
