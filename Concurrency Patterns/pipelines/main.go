package main

import (
	"fmt"
	"time"
)
// multi stage process
func sliceToChan(numbers []int) <-chan int {
	result := make(chan int)
	go func() {
		for _, n := range numbers {
			result <- n
		}
		close(result)
	}()
	return result
}

// processing function
func squareFunc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
//	for measurement
	start := time.Now()
	//input
	nums := []int{1, 4, 5, 6, 2, 8}
	// stage1  
	dataChan := sliceToChan(nums)
	// stage 2
	finalChannel := squareFunc(dataChan)
	//stage 3
	for n := range finalChannel {
		fmt.Println("value is ", n)
	}
	fmt.Println(time.Since(start))
}