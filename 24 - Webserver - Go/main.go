package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler) // pega todos os parametros posterior ao /dir/
	mux.HandleFunc("GET /books/{$}", BooksHandler)            // indica a URL exata (acabando com /books ou /books/)
	mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)
	mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)

	// Essas 2 rotas abaixo não funcionam, elas tem o mesmo grau de precedência, mesmo template.
	// O Go lança um erro de rota ambígua.
	// mux.HandleFunc("GET /books/{id}", BooksPrecedenceHandler)
	// mux.HandleFunc("GET /{x}/latest", BooksPrecedence2Handler)

	http.ListenAndServe(":8080", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d")
	w.Write([]byte(dirpath))
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Book precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	x := r.PathValue("x")
	w.Write([]byte("Book precedence " + x))
}
