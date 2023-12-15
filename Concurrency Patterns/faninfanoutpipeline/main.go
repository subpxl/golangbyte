package main

import (
	"fmt"
	"sync"
	"time"
)

// multi stage process
func sliceToChan(numbers []int) chan int {
	result := make(chan int)
	go func() {
		for _, n := range numbers {
			result <- n
		}
		close(result)
	}()
	return result
}

func squareFunc(wg *sync.WaitGroup, in chan int) chan int {
	defer wg.Done()
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanout(wg *sync.WaitGroup, inchan chan int, n int) []chan int {
	outchans := make([]chan int, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		outchans[i] = squareFunc(wg, inchan)
	}
	return outchans
}

func fanin(wg *sync.WaitGroup, inchans []chan int) chan int {
	outchan := make(chan int)
	go func() {
		wg.Wait()
		for _, ch := range inchans {
			for n := range ch {
				outchan <- n
			}
		}
		close(outchan)

	}()
	return outchan
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	//input
	nums := []int{1, 4, 5, 6, 2, 8}
	// stage1
	dataChan := sliceToChan(nums)

	outChans := fanout(&wg, dataChan, 10)

	finalChannel := fanin(&wg, outChans)
	for n := range finalChannel {
		fmt.Println("value is ", n)
	}
	fmt.Println(time.Since(start))
}
