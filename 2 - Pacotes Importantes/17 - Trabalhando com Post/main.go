package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Teste struct {
	FirstName string `json:"name"`
	LastName  string `json:"last"`
}

func main() {

	var t = Teste{FirstName: "René", LastName: "Oliveira"}

	json, e1 := json.Marshal(t)

	client := http.Client{}

	buf := bytes.NewBuffer([]byte(json))

	resp, e1 := client.Post("http://viacep.com.br/ws/03132010/json/", "application/json", buf)
	panicIfError(e1)

	defer resp.Body.Close()

	fmt.Print(resp.StatusCode)
	fmt.Print(string(json))
}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
