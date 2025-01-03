package main

import "fmt"

type Cliente struct {
	Name  string
	saldo float64
}

func (c Cliente) Simular() float64 {
	return c.saldo * 0.15
}

func (c *Cliente) Efetuar() {
	c.saldo -= c.Simular()

	fmt.Printf("\n%v", &c)
}

func newCliente() *Cliente {
	return &Cliente{Name: "René", saldo: 10.0}
}

func depositar(c *Cliente) {
	c.saldo += 10
}

func main() {

	cliente := *newCliente()

	var simulacao = cliente.Simular()

	fmt.Printf("\nCliente %s saldo %f simulacao %f", cliente.Name, cliente.saldo, simulacao)

	cliente.Efetuar()

	fmt.Printf("\nCliente %s saldo %f", cliente.Name, cliente.saldo)

	depositar(&cliente)

	fmt.Printf("\nCliente %s saldo %f", cliente.Name, cliente.saldo)

}
