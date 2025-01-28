package main

import (
	"fmt"
	"sync"
	"time"
)

// Deadlock Prevention: Proper locking order
func preventDeadlock() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("Goroutine 1: Locked mu1")
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("Goroutine 1: Locked mu2")
	}()

	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("Goroutine 2: Locked mu1")
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("Goroutine 2: Locked mu2")
	}()

	wg.Wait()
}

// Deadlock Detection: Simulating a deadlock
func detectDeadlock() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42 // This write will cause a deadlock as no one reads
	}()

	// Simulate a deadlock (no one is reading from the channel)
	fmt.Println("DetectDeadlock: Waiting for data (this will hang)...")
	time.Sleep(3 * time.Second) // Wait to simulate the issue
}

// Deadlock-Free Design: Using channels for communication
func deadlockFreeDesign() {
	taskChan := make(chan int, 5)

	go func(id int) {
		for task := range taskChan {
			fmt.Printf("Worker %d received task %d\n", id, task)
		}
	}(1)

	go func(id int) {
		for task := range taskChan {
			fmt.Printf("Worker %d received task %d\n", id, task)
		}
	}(2)

	for i := 1; i <= 10; i++ {
		taskChan <- i
	}
	close(taskChan) // Avoid goroutines waiting indefinitely
}

// Deadlock Recovery: Using `select` to timeout
func deadlockRecovery() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42 // Simulate delayed work
	}()

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	case <-time.After(1 * time.Second): // Timeout to recover
		fmt.Println("Timeout! Avoiding deadlock")
	}
}

func main() {
	fmt.Println("=== Deadlock Prevention ===")
	preventDeadlock()

	fmt.Println("\n=== Deadlock Detection (this will hang) ===")
	// Uncomment the next line to see the detection example:
	// detectDeadlock()

	fmt.Println("\n=== Deadlock-Free Design ===")
	deadlockFreeDesign()

	fmt.Println("\n=== Deadlock Recovery ===")
	deadlockRecovery()
}
