package main

import (
	"context"
	"fmt"
	"time"
)

var Counter int

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	done := make(chan bool)

	go Worker(ctx, done)

	// Wait for the worker to complete or timeout
	if <-done {
		fmt.Println("Main: Worker has completed its task.")
	} else {
		fmt.Println("Main: Worker did not complete. timeout or was canceled.")
	}
	fmt.Println("all done now exiting")

}

func Worker(ctx context.Context, done chan bool) {
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("woring good")
			done <- true
		case <-ctx.Done():
			fmt.Println("timeout")
			done <- false
		}
	}
}
