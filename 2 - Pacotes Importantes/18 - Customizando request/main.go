package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Teste struct {
	FirstName string `json:"name"`
	LastName  string `json:"last"`
}

func main() {

	client := http.Client{}

	req, e := http.NewRequest("GET", "https://www.google.com", nil)
	panicIfError(e)

	req.Header.Set("Accept", "application/json")

	resp, e1 := client.Do(req)
	panicIfError(e1)

	defer resp.Body.Close()

	body, e3 := ioutil.ReadAll(resp.Body)
	panicIfError(e3)

	fmt.Print(string(body))

}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
