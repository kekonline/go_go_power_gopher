package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate a long-running operation
func longRunningTask(ctx context.Context, resultChan chan<- string) {
	select {
	case <-time.After(3 * time.Second): // Simulate a delay
		resultChan <- "Task completed"
	case <-ctx.Done():
		fmt.Println("Task canceled:", ctx.Err()) // Handle timeout
	}
}

func main() {
	resultChan := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure we free resources

	go longRunningTask(ctx, resultChan)

	select {
	case result := <-resultChan:
		fmt.Println("Received:", result)
	case <-ctx.Done():
		fmt.Println("Timeout! Task took too long:", ctx.Err())
	}
}
