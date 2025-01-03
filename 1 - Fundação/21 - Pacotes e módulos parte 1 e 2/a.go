package main

import (
	"exercicio21/matematica"
	"fmt"
)

func main() {

	var result = matematica.Soma[int](1, 3)

	fmt.Println("Resultado ", result)

}
