package main

import "fmt"

type ID int

var a = 1.298
var id ID = 999

func main() {
	fmt.Printf("O tipo de a é %T", a)
	fmt.Printf("\nO valor de a é %v", a)
	fmt.Printf("\nO valor de a é %f", a)

	fmt.Printf("\nO tipo de id é %T", id)

}
