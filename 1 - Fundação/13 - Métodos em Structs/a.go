package main

import (
	"fmt"
	"strconv"
)

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

func (c Cliente) Desativar() {
	c.Ativo = false
}

func (e Endereco) ToString() string {
	return e.Logradouro + " " + strconv.Itoa(e.Numero) + " - " + e.Cidade + "/" + e.Estado
}

func main() {

	rene := Cliente{
		Nome:     "René",
		Idade:    46,
		Ativo:    true,
		Address:  Endereco{Logradouro: "Rua Ettore Ximenes"},
		Endereco: Endereco{Logradouro: "Rua Corin", Numero: 147, Cidade: "São Paulo", Estado: "SP"}
	}

	rene.Address.Numero = 10

	rene.Desativar()

	fmt.Printf("Nome %s, idade %d, ativo %t", rene.Nome, rene.Idade, rene.Ativo)
	fmt.Printf("\nEndereço %s, %d", rene.Address.Logradouro, rene.Address.Numero)
	fmt.Printf("\nEndereço %s", rene.Endereco.ToString())
}
