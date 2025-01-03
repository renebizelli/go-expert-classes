package main

import (
	"fmt"
	"strings"
)

func main() {

	linha()

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d - ", i)
	}

	linha()

	values := []string{"1", "2", "3", "4"}

	for index, value := range values {
		fmt.Printf("\n%d: %v", index, value)
	}

	linha()

	var i2 int = 0

	for i2 < 10 {
		fmt.Println(i2)
		i2++
	}

}

func linha() {
	fmt.Println("")
	fmt.Println(strings.Repeat("*", 100))
	fmt.Println("")
}
