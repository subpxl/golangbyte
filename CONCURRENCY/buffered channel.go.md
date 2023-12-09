package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// buffered channel
	myChan := make(chan int,10)

	go Generator(&wg,myChan)

	go func(wg *sync.WaitGroup) {
	defer wg.Done()
		for val := range myChan {
			fmt.Println("value is ", val)
		}
	}(&wg)
	
	wg.Wait()
	fmt.Println("all done now exiting")
}

func Generator(wg *sync.WaitGroup,myChan chan int) {
	defer wg.Done()
	for i := 0; i < 15; i++ {
		myChan <- i
	}
	close(myChan)
}
