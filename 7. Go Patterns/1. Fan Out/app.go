package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker function processes jobs and sends results.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Example result: double the job value
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to the jobs channel
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and display results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
