package main

import (
	"fmt"
	"sync"
)

func main() {

	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publisher(channel)

	go reader(channel, &wg)

	wg.Wait()

}

func publisher(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}

	//close(channel)
}

func reader(channel chan int, wg *sync.WaitGroup) {
	for x := range channel {
		fmt.Println(x)
		wg.Done()
	}
}
