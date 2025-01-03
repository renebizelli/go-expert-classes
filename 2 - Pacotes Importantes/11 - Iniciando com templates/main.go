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

	cursoTemplate := template.New("CursoTemplate")

	tmp, _ := cursoTemplate.Parse("Curso: {{.Nome}}, Carga: {{.Carga}}")

	e := tmp.Execute(os.Stdout, curso)

	if e != nil {
		panic(e)
	}

}
