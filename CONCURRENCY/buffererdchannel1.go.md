package main

import (
	"fmt"
)

func main() {

	// buffered channel
	myChan := make(chan int, 10)

	myChan <- 1
	myChan <- 88
	myChan <- 99

	//	values are present in channel
	//check for values if present
	val, ok := <-myChan
	fmt.Println(val, " ", ok)
	
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)

	fmt.Println("all done now exiting")
}
