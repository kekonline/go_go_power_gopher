package main

import (
	"fmt"
	"time"
)

func slowGreet(phrase string, done chan bool, errChan chan error) {
	defer close(done)
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	if phrase == "" {
		errChan <- fmt.Errorf("empty phrase")
	} else {
		fmt.Println("Hello!", phrase)

	}
}

func greet(phrase string, done chan bool, errChan chan error) {
	if phrase == "" {
		errChan <- fmt.Errorf("empty phrase")
	} else {
		fmt.Println(phrase)
		done <- true
	}
}

func main() {
	done := make(chan bool)
	errChan := make(chan error)

	go greet("Nice to meet you!", done, errChan)
	go greet("How are you?", done, errChan)
	go slowGreet("How ... are ... you ...?", done, errChan)
	go greet("I hope you're liking the course!", done, errChan)

	// wait for all goroutines to finish or an error to occur
	for i := 0; i < 4; i++ {
		select {
		case <-done:
			// A goroutine finished successfully
		case err := <-errChan:
			// An error occurred
			fmt.Println("Error:", err)
			return
		}
	}
}
