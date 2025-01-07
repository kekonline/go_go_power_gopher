package main

import (
	"fmt"
	"time"
)

// Function to demonstrate goroutines
func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from goroutine!")
		time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
	}
}

// Function to demonstrate channels
func countToFive(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send the value to the channel
		time.Sleep(200 * time.Millisecond) // Sleep for 200 milliseconds
	}
	close(ch) // Close the channel when done
}

// Function to demonstrate using select statement with channels
func selectExample(ch1, ch2 <-chan int) {
	for {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Println("Received from channel 1:", val)
			} else {
				fmt.Println("Channel 1 closed")
				return
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Println("Received from channel 2:", val)
			} else {
				fmt.Println("Channel 2 closed")
				return
			}
		}
	}
}

// Main function to run the concurrency examples
func main() {
	// Start a goroutine
	go sayHello()

	// Create a channel for counting
	ch := make(chan int)

	// Start counting in a separate goroutine
	go countToFive(ch)

	// Create another channel for select example
	ch2 := make(chan int)

	// Start another goroutine to send values to the second channel
	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
			time.Sleep(150 * time.Millisecond) // Sleep for 150 milliseconds
		}
		close(ch2) // Close the second channel when done
	}()

	// Use select to receive from both channels
	selectExample(ch, ch2)

	// Use defer to ensure a message is printed at the end
	defer fmt.Println("Finished executing concurrency examples.")
}