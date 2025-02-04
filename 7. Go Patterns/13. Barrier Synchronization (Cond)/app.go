package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 3

func worker(id int, cond *sync.Cond, counter *int, mutex *sync.Mutex) {
	fmt.Printf("Worker %d: Doing some work...\n", id)
	time.Sleep(time.Duration(id) * time.Second) // Simulate work

	mutex.Lock()
	*counter++
	if *counter == numWorkers {
		fmt.Println("All workers reached the barrier. Broadcasting signal...")
		cond.Broadcast() // Wake up all waiting goroutines
	} else {
		fmt.Printf("Worker %d: Waiting at the barrier...\n", id)
		cond.Wait() // Wait for the signal
	}
	mutex.Unlock()

	fmt.Printf("Worker %d: Passed the barrier and continuing...\n", id)
}

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	counter := 0

	for i := 1; i <= numWorkers; i++ {
		go worker(i, cond, &counter, &mutex)
	}

	time.Sleep(5 * time.Second) // Ensure main doesn't exit early
}
