package main

import "fmt"

func changeStatus(status *string) {

	if *status == "open" {
		*status = "close"
	} else {
		*status = "open"
	}
}

func main() {

	var status string = "open"

	fmt.Printf("\nStatus %s", status)

	changeStatus(&status)

	fmt.Printf("\nStatus %s", status)

	changeStatus(&status)

	fmt.Printf("\nStatus %s", status)
}
