package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome  string
	Carga int
}

func main() {

	curso := []Curso{
		{Nome: "GO", Carga: 40},
		{Nome: "C#", Carga: 45},
	}

	t := template.Must(template.New("template.html").ParseFiles("./public/template.html"))

	e := t.Execute(os.Stdout, curso)

	if e != nil {
		panic((e))
	}

}
