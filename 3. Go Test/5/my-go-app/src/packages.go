// packages.go
// This file explains how to organize code into packages, export and import identifiers, and use third-party packages.

// Package declaration
package main

// Importing necessary packages
import (
    "fmt" // Standard library for formatted I/O
    "math/rand" // Standard library for generating random numbers
    "time" // Standard library for time-related functions
)

// Function to demonstrate package usage
func demonstratePackages() {
    // Using the fmt package to print a message
    fmt.Println("Demonstrating package usage in Go!")

    // Using the time package to get the current time
    currentTime := time.Now()
    fmt.Println("Current time:", currentTime)

    // Using the rand package to generate a random number
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    randomNum := rand.Intn(100) // Generate a random number between 0 and 99
    fmt.Println("Random number:", randomNum)
}

// The main function serves as the entry point of the application
func main() {
    // Call the function to demonstrate package usage
    demonstratePackages()
}