package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

// main é a thread 1
func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	go task("A", &waitGroup) //thread 2
	go task("B", &waitGroup) //thread 3
	go func() {              //thread 4
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: anonymous task is running\n", i)
			time.Sleep(1 * time.Second)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
