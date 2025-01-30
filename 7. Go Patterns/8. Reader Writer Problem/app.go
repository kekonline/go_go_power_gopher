package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data  []int        // Shared resource
	mutex sync.RWMutex // Read-Write Mutex
)

func writeData(id int, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() // Lock for writing
	fmt.Printf("Writer %d is writing value %d\n", id, value)
	data = append(data, value)
	time.Sleep(500 * time.Millisecond) // Simulate a time-consuming write operation
	mutex.Unlock()                     // Unlock after writing
	fmt.Printf("Writer %d finished writing value %d\n", id, value)
}

func readData(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.RLock() // Lock for reading
	fmt.Printf("Reader %d is reading data: %v\n", id, data)
	time.Sleep(200 * time.Millisecond) // Simulate a time-consuming read operation
	mutex.RUnlock()                    // Unlock after reading
	fmt.Printf("Reader %d finished reading\n", id)
}

func main() {
	var wg sync.WaitGroup
	data = make([]int, 0) // Initialize the shared slice

	// Start some writer goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go writeData(i, i*10, &wg)
	}

	// Start some reader goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go readData(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines completed. Final data:", data)
}
