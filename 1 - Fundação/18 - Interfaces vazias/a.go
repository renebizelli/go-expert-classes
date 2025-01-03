package main

import "fmt"

type i1 interface{}

type Cliente struct {
	Nome string
}

func (c Cliente) Printar() {
	fmt.Println(c.Nome)
}

type Contato struct {
	Nome string
}

func (c Contato) Printar() {
	fmt.Println(c.Nome)
}

func Print(o i1) {
	fmt.Printf("o: %v %T\n", o, o)
}

func main() {

	cl := Cliente{Nome: "AAaaaaa"}
	co := Contato{Nome: "cccccc"}

	Print(cl)
	Print(co)

}
