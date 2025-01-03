package main

import "fmt"

func main() {

	channel := make(chan int)

	go publisher(channel)

	reader(channel)

}

func publisher(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}

	close(channel)
}

func reader(channel chan int) {
	for x := range channel {
		fmt.Println(x)
	}
}
