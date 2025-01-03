package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
	Endereco
}

func main() {

	rene := Cliente{
		Nome:     "René",
		Idade:    46,
		Ativo:    true,
		Address:  Endereco{Logradouro: "Rua Ettore Ximenes"},
		Endereco: Endereco{Logradouro: "Rua Corin"},
	}

	rene.Address.Numero = 10
	rene.Endereco.Numero = 147

	rene.Ativo = false

	fmt.Printf("Nome %s, idade %d, ativo %t", rene.Nome, rene.Idade, rene.Ativo)
	fmt.Printf("\nEndereço %s, %d", rene.Address.Logradouro, rene.Address.Numero)
	fmt.Printf("\nEndereço %s, %d", rene.Endereco.Logradouro, rene.Endereco.Numero)
}
