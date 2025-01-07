// file_operations.go
package main

import (
	"fmt"
	"io/ioutil" // Package for reading files
	"os"        // Package for file operations
)

// Function to write data to a file
func writeFile(filename string, data string) {
	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		// Handle error if file creation fails
		panic(err) // Stop execution and print the error
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write data to the file
	_, err = file.WriteString(data)
	if err != nil {
		// Handle error if writing fails
		panic(err) // Stop execution and print the error
	}
	fmt.Println("Data written to file:", filename)
}

// Function to read data from a file
func readFile(filename string) {
	// Read the content of the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// Handle error if reading fails
		panic(err) // Stop execution and print the error
	}
	fmt.Println("Data read from file:", string(data)) // Print the content of the file
}

// Main function to demonstrate file operations
func main() {
	// Define the filename
	filename := "example.txt"

	// Data to write to the file
	data := "Hello, this is a test file."

	// Write data to the file
	writeFile(filename, data)

	// Read data from the file
	readFile(filename)
}