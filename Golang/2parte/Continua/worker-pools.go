package main

import (
	"fmt"
	"time"
)

// WorkerPool is a struct that represents a pool of workers.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		// Send the result back to the results channel
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// Start a fixed number of workers

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// Send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel to signal no more jobs

	for a := 1; a <= numJobs; a++ {
		<-results // Receive results from the workers
	}
}