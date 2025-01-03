package main

import (
	"fmt"

	"github.com/renebizelli/goexpert/5-Packaging/1/math"
)

func main() {

	math := math.Math{A: 1, B: 2}
	fmt.Println(math.Add())

}
