package main

import "fmt"

func main() {

	var a = map[string]int{"b": 2}

	a["a"] = 1
	a["c"] = 3

	fmt.Println(a["a"])
	fmt.Println(a["b"])
	fmt.Println(a["c"])

	delete(a, "c")

	fmt.Println(a["c"])

	for key, value := range a {
		fmt.Printf("%s %d\n", key, value)
	}

	for _, value := range a {
		fmt.Printf("%d\n", value)
	}

	for key := range a {
		fmt.Printf("%s\n", key)
	}

}
