package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	for i := 0; i < 100; i++ {
		go worker(i, ch)
	}

	for j := 0; j < 100; j++ {
		ch <- j
	}

	// go worker(9999999, ch)
	// ch <- 0
}

func worker(id int, channel <-chan int) {

	for x := range channel {
		fmt.Printf("Worker %d received %d\n", id, x)
		time.Sleep(time.Second)
	}

}
