package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
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

func AddH(value int) string {
	return fmt.Sprintf("%dh", value)
}

func handler(w http.ResponseWriter, r *http.Request) {
	curso := []Curso{
		{Nome: "GO", Carga: 40},
		{Nome: "C#", Carga: 45},
	}

	ts := []string{"./public/template.html", "./public/content.html"}

	t := template.New("template.html")
	t.Funcs(template.FuncMap{
		"AddH":  AddH,
		"Upper": strings.ToUpper})

	t = template.Must(t.ParseFiles(ts...))

	e := t.Execute(w, curso)

	if e != nil {
		panic((e))
	}
}
