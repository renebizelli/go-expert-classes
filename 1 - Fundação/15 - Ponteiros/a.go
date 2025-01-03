package main

import "fmt"

func main() {

	var a int = 100
	fmt.Printf("valor de a: %d", a)

	var ponteiro_a *int = &a
	fmt.Printf("\nponteiro de a: %v", ponteiro_a)

	*ponteiro_a = 200
	fmt.Printf("\nvalor de a: %d", a)
	fmt.Printf("\nvalor de a pelo ponteiro: %d", *ponteiro_a)
	fmt.Printf("\nvalor de a pelo ponteiro: %v", *ponteiro_a)
	fmt.Printf("\ntipo do valor de a pelo ponteiro: %T", *ponteiro_a)
	fmt.Printf("\ntipo do ponteiro: %T", ponteiro_a)

	*&a = 300
	fmt.Printf("\nvalor de a: %d", a)
}
