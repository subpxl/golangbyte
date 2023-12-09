package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup 
	//	add number of goroutines to track for
	wg.Add(1)
	//	run goroutine with wg as parameter
	go MyFunc(&wg)
	//	wait for all goroutines to finish
	wg.Wait()

	fmt.Println("all done now exiting")
}

func MyFunc(wg *sync.WaitGroup) {
//	call wg.done after function is executed
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("hello index is : ", i)
	}
}
