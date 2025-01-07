// error_handling.go
// This file covers basic error handling in Go, including panic and recover.
// It shows how to handle errors gracefully.

package main

import (
	"fmt"
	"os"
)

// function that may cause an error
func riskyOperation() {
	// Simulating a panic situation
	panic("Something went wrong!")
}

// function to demonstrate recover
func safeOperation() {
	// Using defer to ensure that recover is called
	defer func() {
		if r := recover(); r != nil {
			// Recovering from panic and printing the error
			fmt.Println("Recovered from:", r)
		}
	}()
	
	// Calling a function that may panic
	riskyOperation()
	
	// This line will not execute if panic occurs
	fmt.Println("This line will not be printed if panic happens.")
}

func main() {
	// Calling the safe operation to demonstrate error handling
	fmt.Println("Starting the application...")
	safeOperation()
	fmt.Println("Application finished successfully.")
	
	// Example of handling file errors
	file, err := os.Open("non_existent_file.txt")
	if err != nil {
		// Handling the error gracefully
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after operations

	// If we reach here, the file opened successfully
	fmt.Println("File opened successfully.")
}