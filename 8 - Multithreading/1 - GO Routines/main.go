package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}

}

// main é a thread 1
func main() {
	go task("A") //thread 2
	go task("B") //thread 3

	time.Sleep(15 * time.Second)
}
