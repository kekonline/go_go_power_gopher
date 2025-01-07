package main

import "fmt"

// This file demonstrates the use of variables, values, and operators in Go.

// Declare a global variable
var globalVar int = 10 // This is a global variable of type int

// Function to demonstrate variable usage
func variableDemo() {
    // Local variable declaration
    var localVar int = 5 // This is a local variable of type int

    // Demonstrating value types
    fmt.Println("Global Variable:", globalVar) // Print the global variable
    fmt.Println("Local Variable:", localVar)    // Print the local variable

    // Demonstrating null values
    var nullVar *int // This is a pointer variable that is currently nil (null)
    fmt.Println("Null Variable:", nullVar)      // Print the null variable

    // Demonstrating type conversions
    var floatVar float64 = 3.14 // A float64 variable
    intVar := int(floatVar)      // Convert float64 to int
    fmt.Println("Converted Integer:", intVar) // Print the converted integer

    // Demonstrating operators
    sum := globalVar + localVar // Addition operator
    fmt.Println("Sum of Global and Local Variable:", sum) // Print the sum

    // Using a simple operator
    isEqual := (globalVar == localVar) // Comparison operator
    fmt.Println("Are Global and Local Variables Equal?", isEqual) // Print the comparison result
}

// The main function is the entry point of the application
func main() {
    variableDemo() // Call the variable demonstration function
}