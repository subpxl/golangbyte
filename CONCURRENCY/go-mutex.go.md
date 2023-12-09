package main

import (
	"fmt"
	"sync"
)

var Counter int

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	go Increment(&wg, &mu)
	go Decrement(&wg, &mu)

	wg.Wait()
	fmt.Println("counter is :", Counter)
	fmt.Println("all done now exiting")
}

func Increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		Counter ++
		mu.Unlock()
	}
}
func Decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		Counter--
		mu.Unlock()
	}
}
