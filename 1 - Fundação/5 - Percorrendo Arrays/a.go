package main

import "fmt"

func main() {

	var a [3]int
	a[0] = 1

	fmt.Printf("len >> %v", len(a))
	fmt.Printf("\n>> %v", a[0])

	for i, v := range a {
		fmt.Printf("\n Valor do índice: %d, valor %v", i, v)
	}

	for i := range a {
		fmt.Printf("\n >> %d", i)
	}

	for i := 0; i < len(a)-1; i++ {
		fmt.Printf("\n valor do índice: %d, valor %v", i, a[i])

	}

}
