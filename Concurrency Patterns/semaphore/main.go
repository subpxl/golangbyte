package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, size)}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}

func main() {
	start := time.Now()
	numWorkers := 10
	numArr := []int{5, 7, 4, 77, 2, 88, 66, 97, 90, 45, 34, 12, 78}
	var wg sync.WaitGroup
	sem := NewSemaphore(3)
	jobsChan := make(chan int, len(numArr))
	results := make(chan int, len(numArr))
	wg.Add(numWorkers)
	// Start workers
	worker(numWorkers, &wg, jobsChan, results, sem)
	// Add job to process in jobs channel
	for _, job := range numArr {
		sem.Acquire()
		jobsChan <- job
	}
	close(jobsChan)
	// Wait for all workers to finish
	wg.Wait()
	// Close the results channel after all jobs are done
	close(results)
	// Read results from the channel
	for res := range results {
		fmt.Println("Result:", res)
	}
	fmt.Println(time.Since(start))
}

func worker(numWorkers int, wg *sync.WaitGroup, jobs chan int, results chan int, sem *Semaphore) {
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Println("Worker", workerID, "started job", job)
				time.Sleep(1 * time.Second)
				results <- job * 10
				fmt.Println("Worker", workerID, "finished job", job)
				sem.Release()
			}
		}(i)
	}
}
