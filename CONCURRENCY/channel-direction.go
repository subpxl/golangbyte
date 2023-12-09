package main

import (
	"fmt"
	"time"
)

func main() {
//	example sendonly channel
	sendOnly := make(chan<- int, 5)

//	example recieveonly channel
	receiveOnly := make(<-chan int)

	biDirectional := make(chan int)

	go func(ch chan<- int) {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}(biDirectional)

	go func(ch <-chan int) {
		for val := range ch {
			fmt.Println("Received value:", val)
		}
	}(biDirectional)

	fmt.Println("done execution")

}
