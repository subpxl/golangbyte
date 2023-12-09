package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	var wg sync.WaitGroup
	const goroutines = 100

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			atomic.AddInt64(&counter,1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", atomic.LoadInt64(&counter)) // Read the final counter value atomically
}
