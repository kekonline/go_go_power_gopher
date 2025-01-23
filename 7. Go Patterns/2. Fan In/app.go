package main

import (
	"fmt"
	"time"
)

// Function that simulates a worker sending data to a channel
func worker(id int, c chan string) {
	for i := 0; i < 3; i++ {
		c <- fmt.Sprintf("Worker %d: Message %d", id, i)
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
	close(c)
}

// Function to fan-in multiple channels into one
func fanIn(channels ...chan string) chan string {
	out := make(chan string)

	go func() {
		for _, ch := range channels {
			for msg := range ch {
				out <- msg
			}
		}
		close(out)
	}()

	return out
}

func main() {
	// Create channels for workers
	worker1 := make(chan string)
	worker2 := make(chan string)
	worker3 := make(chan string)

	// Start workers
	go worker(1, worker1)
	go worker(2, worker2)
	go worker(3, worker3)

	// Use fanIn to merge channels
	merged := fanIn(worker1, worker2, worker3)

	// Read from the merged channel
	for msg := range merged {
		fmt.Print("Received: ")
		fmt.Println(msg)
	}

	fmt.Println("All messages received!")
}
