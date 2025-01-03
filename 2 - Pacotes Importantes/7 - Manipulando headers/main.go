package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/cep/{cep}", BuscaCepHandler)
	http.HandleFunc("/cep/", BuscaCepHandler)

	http.ListenAndServe(":8080", nil)

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	cep := r.PathValue("cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(r.PathValue("cep")))
}
