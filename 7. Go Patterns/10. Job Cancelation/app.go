package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start the worker goroutine
	go worker(ctx)

	// Simulate some work in main
	time.Sleep(2 * time.Second)
	fmt.Println("Main: Canceling worker...")

	// Cancel the worker
	cancel()

	// Give some time to see the cancellation effect
	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Received cancellation signal")
			return
		default:
			fmt.Println("Worker: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
