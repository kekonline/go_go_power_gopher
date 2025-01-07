// debugging_testing.go
// This file provides guidance on debugging and testing Go programs, including best practices and tools.

package main

import (
	"fmt"
	"log"
	"os"
)

// Debugging function to demonstrate logging
func debugExample() {
	// Log a debug message
	log.Println("Debugging example function called")
}

// Testing function to demonstrate a simple test case
func testExample() {
	// Simulate a test case
	expected := 5
	result := add(2, 3)

	// Check if the result matches the expected value
	if result != expected {
		log.Fatalf("Test failed: expected %d, got %d", expected, result)
	} else {
		log.Println("Test passed!")
	}
}

// Simple function to add two numbers
func add(a int, b int) int {
	return a + b
}

// Main function to run debugging and testing examples
func main() {
	// Set up logging to a file
	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	log.SetOutput(file)

	// Call the debugging example
	debugExample()

	// Call the testing example
	testExample()
}