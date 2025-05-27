package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	ctx := context.Background()

	fmt.Println(url.QueryEscape("Rio de Janeiro"))

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://api.weatherapi.com/v1/current.json?q=%s", url.QueryEscape("Rio de Janeiro")), nil)

	if err != nil {
		fmt.Printf("ERROOO0 %s", err.Error())

		panic(err)
	}

	req.Header.Set("Key", "b0521787020347eca0b22548250104")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("ERROOO1 %s", err.Error())

		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("ERROOO2 %s", err.Error())
		panic(err)
	}

	fmt.Printf("body %s", string(body))

}
