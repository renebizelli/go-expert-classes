package main

import "fmt"

func main() {

	var r int = sum(1, 2, 3, 4, 5)

	fmt.Printf("Sum %d", r)
}

func sum(n ...int) int {

	total := 0

	for _, v := range n {
		total += v
	}

	return total
}
