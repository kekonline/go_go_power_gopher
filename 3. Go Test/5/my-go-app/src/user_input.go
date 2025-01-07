// user_input.go
package main

import (
	"fmt" // Importing the fmt package for formatted I/O
)

// Function to handle user input
func handleUserInput() {
	var name string // Declare a variable to store the user's name
	var age int    // Declare a variable to store the user's age

	// Prompt the user for their name
	fmt.Print("Enter your name: ")
	// Read the user's input and store it in the name variable
	_, err := fmt.Scan(&name) // Using fmt.Scan() to read input
	if err != nil { // Check for errors during input
		panic("Error reading name") // Panic if there's an error
	}

	// Prompt the user for their age
	fmt.Print("Enter your age: ")
	// Read the user's input and store it in the age variable
	_, err = fmt.Scan(&age) // Using fmt.Scan() to read input
	if err != nil { // Check for errors during input
		panic("Error reading age") // Panic if there's an error
	}

	// Output the user's name and age
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age) // Print formatted output
}

// The main function is the entry point of the application
func main() {
	handleUserInput() // Call the function to handle user input
}