package main

import "fmt"

type i1 interface{}

type Cliente struct {
	Nome string
}

func (c Cliente) PrintarCliente() {
	fmt.Println(c.Nome)
}

type Contato struct {
	Nome string
}

func (c Contato) PrintarContato() {
	fmt.Println(c.Nome)
}

func Print(o i1) {

	cl, eCl := o.(Cliente)
	co, eCo := o.(Contato)

	if eCl {
		cl.PrintarCliente()
	} else if eCo {
		co.PrintarContato()
	}

	fmt.Printf("o: %v %T\n", o, o)
}

func main() {

	cl := Cliente{Nome: "AAaaaaa"}
	co := Contato{Nome: "cccccc"}

	Print(cl)
	Print(co)

}
