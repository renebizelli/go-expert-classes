package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome  string
	Carga int
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	http.ListenAndServe(":8080", mux)

}

func handler(w http.ResponseWriter, r *http.Request) {
	curso := []Curso{
		{Nome: "GO", Carga: 40},
		{Nome: "C#", Carga: 45},
	}

	t := template.Must(template.New("template.html").ParseFiles("./public/template.html"))

	e := t.Execute(w, curso)

	if e != nil {
		panic((e))
	}
}
