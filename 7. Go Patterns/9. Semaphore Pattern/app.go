package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxWorkers = 3                         // Limit concurrency to 3 workers
	semaphore := make(chan struct{}, maxWorkers) // Buffered channel as semaphore
	var wg sync.WaitGroup

	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, task := range tasks {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire semaphore

		go func(task int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release semaphore

			fmt.Printf("Processing task %d\n", task)
			time.Sleep(2 * time.Second) // Simulate work
		}(task)
	}

	wg.Wait()
	fmt.Println("All tasks completed!")
}
