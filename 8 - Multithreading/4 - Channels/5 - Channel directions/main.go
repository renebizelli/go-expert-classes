package main

import (
	"fmt"
)

func main() {

	channel := make(chan string)

	go recebe("rene bizelli", channel)

	envia(channel)
}

func recebe(texto string, channel chan<- string) {
	channel <- texto
}

func envia(channel <-chan string) {
	fmt.Println(<-channel)
}
