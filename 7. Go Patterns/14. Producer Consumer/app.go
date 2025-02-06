package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function
func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ { // Each producer generates 5 jobs
		job := rand.Intn(100) // Random job
		fmt.Printf("Producer %d: produced job %d\n", id, job)
		jobs <- job
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

// Consumer function
func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Consumer %d: processing job %d\n", id, job)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan int, 10) // Buffered channel
	var wg sync.WaitGroup

	// Start producers
	numProducers := 3
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i, jobs, &wg)
	}

	// Start consumers
	numConsumers := 2
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, jobs, &wg)
	}

	// Wait for all producers to finish
	wg.Wait()
	close(jobs) // Close channel when producers finish

	// Wait for remaining consumers to finish
	wg.Add(numConsumers)
	for i := 1; i <= numConsumers; i++ {
		go consumer(i, jobs, &wg)
	}

	wg.Wait() // Ensure all consumers finish before exiting
}
