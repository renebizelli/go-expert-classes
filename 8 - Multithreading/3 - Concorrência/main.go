package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count uint32 = 0

func main() {

	l := 5
	q := 10000

	wg := sync.WaitGroup{}
	wg.Add(l)

	for i := 0; i < l; i++ {
		go task(i, q, &wg)
	}

	wg.Wait()

	fmt.Println(count)
}

func task(index, q int, wg *sync.WaitGroup) {

	fmt.Printf("Start %d\n", index)

	for i := 0; i < q; i++ {
		atomic.AddUint32(&count, 1)
		fmt.Printf("Thread %d: %d - %d\n", index, i, count)
		time.Sleep(100 * time.Microsecond)
	}

	wg.Done()
}
