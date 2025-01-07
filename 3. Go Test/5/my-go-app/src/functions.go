// functions.go
package main

import "fmt"

// Function that takes a variable number of arguments and returns their sum
func sum(numbers ...int) int {
    total := 0 // Initialize total to 0
    for _, number := range numbers { // Loop through each number
        total += number // Add each number to total
    }
    return total // Return the total sum
}

// Function that returns a closure that increments a counter
func counter() func() int {
    count := 0 // Initialize count to 0
    return func() int { // Return an anonymous function
        count++ // Increment count
        return count // Return the current count
    }
}

// Function that demonstrates the use of an anonymous function
func demonstrateAnonymousFunction() {
    // Define an anonymous function that prints a message
    messagePrinter := func(message string) {
        fmt.Println(message) // Print the message
    }
    messagePrinter("Hello from the anonymous function!") // Call the anonymous function
}

// Variadic function that takes a string and prints it multiple times
func repeatString(times int, str string) {
    for i := 0; i < times; i++ { // Loop for the specified number of times
        fmt.Println(str) // Print the string
    }
}

// Main function to demonstrate the usage of the above functions
func main() {
    // Demonstrate the sum function
    total := sum(1, 2, 3, 4, 5) // Call sum with multiple arguments
    fmt.Println("Total Sum:", total) // Print the total sum

    // Demonstrate the counter function
    increment := counter() // Get the closure
    fmt.Println("Counter:", increment()) // Call the closure to get the first count
    fmt.Println("Counter:", increment()) // Call again to get the second count

    // Demonstrate the anonymous function
    demonstrateAnonymousFunction() // Call the function to show the message

    // Demonstrate the variadic function
    repeatString(3, "Hello, Go!") // Call the function to repeat the string
}