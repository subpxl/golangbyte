package main

import (
	"fmt"
)

func main() {
	myChan := make(chan int)
	done := make(chan bool)
	
	go Increment(myChan,done)
	for {
		select {
		case msg1 :=  <-myChan:
			fmt.Println("value is ", msg1)
		case <-done:
			fmt.Println("exiting ")
			return
		default:
			fmt.Println("no activity")
		}
	}

}

func Increment(myChan chan int,done chan bool) {
	for i := 0; i < 10; i++ {
		myChan <- 99
	}
	done <- true
}
