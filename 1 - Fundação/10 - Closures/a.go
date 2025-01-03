package main

import "fmt"

func main() {

	sum := func(n ...int) int {

		total := 0

		for _, v := range n {
			total += v
		}

		return total
	}

	total := func(a, b, c int) int {
		return sum(a, b, c)
	}(1, 2, 3)

	fmt.Printf("Sum %d", total)

}
