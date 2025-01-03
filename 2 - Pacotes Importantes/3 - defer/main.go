package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	simpleExempla()

	httpExample()

}

func simpleExempla() {

	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
	defer fmt.Println("Quarta linha")
	fmt.Println("Quinta linha")
	defer fmt.Println("Sexta linha")

}

func httpExample() {
	req, err := http.Get("https://www.google.com")

	fmt.Printf("Status code %v", req.StatusCode)

	panicIfError(err)

	defer req.Body.Close() // retarda o fechamento, fecha somente no final

	res, err := io.ReadAll(req.Body)

	panicIfError(err)

	fmt.Print(string(res)[0:100])

	fmt.Printf("Status code %v", req.StatusCode)
}

func panicIfError(err error) {
	if err != nil {
		fmt.Println("Ocorreu algum erro ")
		panic(err)
	}
}
