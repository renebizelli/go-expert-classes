package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}

	}()

	doPanic()

}

func doPanic() {
	panic("Some error ...")
}
