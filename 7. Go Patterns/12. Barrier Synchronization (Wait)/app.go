package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 3

func worker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done()

	// Simulate some work
	fmt.Printf("Worker %d: Doing some work...\n", id)
	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("Worker %d: Waiting at the barrier...\n", id)
	barrier.Done() // Mark this worker as reached the barrier
	barrier.Wait() // Wait for all workers to reach the barrier

	fmt.Printf("Worker %d: Passed the barrier and continuing...\n", id)
}

func main() {
	var wg sync.WaitGroup
	var barrier sync.WaitGroup
	wg.Add(numWorkers)
	barrier.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, &barrier)
	}

	wg.Wait()
	fmt.Println("All workers have finished execution.")
}
