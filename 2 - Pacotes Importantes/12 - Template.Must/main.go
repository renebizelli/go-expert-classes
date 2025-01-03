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

	curso := Curso{Nome: "GO", Carga: 40}

	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}}, Carga: {{.Carga}}"))

	e := t.Execute(os.Stdout, curso)

	if e != nil {
		panic(e)
	}

}
