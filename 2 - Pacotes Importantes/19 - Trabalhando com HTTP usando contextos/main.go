package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)

	defer cancel()

	req, e := http.NewRequestWithContext(ctx, "GET", "https://www.google.com", nil)
	panicIfError(e)

	req.Header.Set("Accept", "application/json")

	resp, e1 := http.DefaultClient.Do(req)
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
