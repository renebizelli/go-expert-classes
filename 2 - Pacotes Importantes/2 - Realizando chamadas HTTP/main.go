package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://www.google.com")

	fmt.Printf("Status code %v", req.StatusCode)

	panicIfError(err)

	res, err := io.ReadAll(req.Body)

	panicIfError(err)

	fmt.Print(string(res))

	fmt.Printf("Status code %v", req.StatusCode)

}

func panicIfError(err error) {
	if err != nil {
		fmt.Println("Ocorreu algum erro ")
		panic(err)
	}
}
