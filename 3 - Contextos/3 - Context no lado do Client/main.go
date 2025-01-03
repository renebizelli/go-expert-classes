package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	// rodar server da aula anterior

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	req, e1 := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	panicIfError(e1)

	res, e2 := http.DefaultClient.Do(req)
	panicIfError(e2)

	defer res.Body.Close()

	body, e3 := io.ReadAll(res.Body)
	panicIfError(e3)

	fmt.Printf(string(body))
}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
