package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	//	to process this
	numArr := []int{5, 7, 4, 77, 2}
	jobs := make(chan int, len(numArr))
	results := make(chan int, len(numArr))
	//	prepare waitgroups
	wg.Add(numWorkers)
	//	start workers with jobs
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, results)
	}
	//	add data to process in jobs channel
	for _, i := range numArr {
		jobs <- i
	}
	close(jobs)
	//	read result from result channel
	for i := 0; i < len(numArr); i++ {
		fmt.Println(<-results)
	}
}

func worker(id int, jobs chan int, results chan int) {
	for job := range jobs {
		fmt.Println("worker ", id, "started job ", job)
		time.Sleep(1 * time.Second)
		results <- job * 10
		fmt.Println("worker ", id, "finished job ", job)
	}
}
