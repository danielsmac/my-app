package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	// Simulate work
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine completes
			worker(i)
		}()
	}
	wg.Wait() // Wait for all goroutines to finish

}