package main

import (
	"fmt"
	"time"
)

// Simulate a task that performs work and can be quit using a quit signal
func taskWithQuitSignal(id int, quitChan chan bool, doneChan chan bool) {
	for {
		select {
		case <-quitChan:
			// If quit signal is received, stop the task
			fmt.Printf("Task %d stopped.\n", id)
			doneChan <- true // Notify that the task is done
			return
		default:
			// Simulate some work
			fmt.Printf("Task %d is running...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create channels for quitting tasks and signaling completion
	quitChannel := make(chan bool)
	doneChannel := make(chan bool, 3) // Buffered channel to handle multiple done signals

	// Launch multiple goroutines for different tasks
	for i := 1; i <= 3; i++ {
		go taskWithQuitSignal(i, quitChannel, doneChannel)
	}

	// Simulate some runtime of the tasks
	time.Sleep(5 * time.Second)

	// Send quit signal to stop all tasks
	close(quitChannel)

	// Wait for all tasks to finish by receiving done signals
	for i := 1; i <= 3; i++ {
		<-doneChannel
	}

	// Print final message
	fmt.Println("All tasks are done, exiting application.")
}
