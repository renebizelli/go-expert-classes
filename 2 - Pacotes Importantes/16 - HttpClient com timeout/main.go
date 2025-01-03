package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	client := http.Client{Timeout: time.Second}

	resp, e1 := client.Get("http://viacep.com.br/ws/03132010/json/")
	panicIfError(e1)

	defer resp.Body.Close()

	body, e2 := io.ReadAll(resp.Body)
	panicIfError(e2)

	fmt.Print(string(body))

}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
