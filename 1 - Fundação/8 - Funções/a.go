package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(sum(1, 3))

	var x, y = sum2(4, 3)

	fmt.Println(x)
	fmt.Println(y)

	var x1, e1 = sumError(4, 3)

	fmt.Println(x1)
	fmt.Println(e1)
	fmt.Printf("%T", e1)

}

func sum(a, b int) int {
	return a + b
}

func sum2(a, b int) (int, bool) {
	return (a + b), a > b
}

func sumError(a, b int) (int, error) {

	var e error = nil

	if a > b {
		e = errors.New("Invalid")
	}

	return a + b, e
}
