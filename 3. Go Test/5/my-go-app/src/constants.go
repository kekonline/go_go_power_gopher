// constants.go
package main

import "fmt"

// Define constants for the application
const (
    Pi       = 3.14 // Constant for the value of Pi
    MaxUsers = 100   // Constant for the maximum number of users
)

// Define a constant for a greeting message
const GreetingMessage = "Welcome to My Go App!"

// Function to demonstrate the use of constants
func demonstrateConstants() {
    // Print the constants
    fmt.Println("Value of Pi:", Pi)
    fmt.Println("Max Users Allowed:", MaxUsers)
    fmt.Println(GreetingMessage)

    // Example of type conversion
    var radius float64 = 5.0
    area := Pi * radius * radius // Calculate area using the constant Pi
    fmt.Printf("Area of circle with radius %.2f: %.2f\n", radius, area)
}