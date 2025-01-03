package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Executando a rotina", i)
		}

		forever <- "Fim da execução"
	}()

	<-forever
}
