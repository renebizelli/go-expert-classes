package main

import (
	"context"
	"fmt"
	"os"
)

func main() {

	// rodar server da aula anterior

	ctx := context.WithValue(context.Background(), "token", "12334343")

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	fmt.Print(ctx.Value("token"))
}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
