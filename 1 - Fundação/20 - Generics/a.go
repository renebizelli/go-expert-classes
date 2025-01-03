package main

import "fmt"

type Guitarra struct {
	Cordas int
}

type Bateria struct {
	Tambores int
}

type Constraint interface {
	Guitarra | Bateria
}

type Instrumento[T1, T2 Constraint] struct {
	I1 T1
	I2 T2
}

func main() {

	var banda = Instrumento[Guitarra, Bateria]{I1: Guitarra{Cordas: 6}, I2: Bateria{Tambores: 4}}

	fmt.Println(banda.I1.Cordas)
	fmt.Println(banda.I2.Tambores)

}
