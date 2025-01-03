package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	rene := Cliente{
		Nome:  "René",
		Idade: 46,
		Ativo: true,
	}

	rene.Ativo = false

	fmt.Printf("Nome %s, idade %d, ativo %t", rene.Nome, rene.Idade, rene.Ativo)
}
