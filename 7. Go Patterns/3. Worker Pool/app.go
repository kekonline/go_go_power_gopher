package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID int
}

// Worker processes tasks from the jobs channel and sends results to the results channel.
func Worker(id int, jobs <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started task %d\n", id, job.ID)
		time.Sleep(1 * time.Second) // Simulating task processing
		fmt.Printf("Worker %d finished task %d\n", id, job.ID)
		results <- fmt.Sprintf("Task %d processed by worker %d", job.ID, id)
	}
}

func main() {
	numJobs := 10   // Total number of tasks
	numWorkers := 3 // Number of workers
	jobs := make(chan Task, numJobs)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	// Send jobs to the workers
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Task{ID: j}
		}
		close(jobs) // Close the jobs channel to signal workers there are no more tasks
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect and display results
	for result := range results {
		fmt.Println(result)
	}
}
